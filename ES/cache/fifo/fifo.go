package fifo

import (
	"container/list"
	"fmt"
	"runtime"
)

type Value interface {
	Len() int
}

type fifo struct {
	//缓存最大的容量，单位是字节
	maxBytes int
	//当一个entry从缓存中移除该回调函数时，默认为nil。
	//groupcache 中的任意可比较类型，value是interface{}
	onEvicted func(key string, value interface{})

	//已使用的字节数，只包括值，key不算
	usedBytes int
	ll        *list.List
	cache     map[string]*list.Element
}
type entry struct {
	key   string
	value interface{}
}

func (e *entry) Len() int {
	return CalcLen(e.value)

}

func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, uint8, int8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.value", value))
	}
	return n
}

//创建一个新的cache，如果maxBytes=0，表示没有容量限制
func New(maxBytes int, onEvicted func(key string, value interface{})) *fifo {
	return &fifo{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}

//set往cache尾部增加一个元素（如果存在，则放入尾部，并修改值）
func (f *fifo) Set(key string, value interface{}) {
	if e, ok := f.cache[key]; ok {
		f.ll.MoveToBack(e)
		en := e.Value.(*entry)
		f.usedBytes = f.usedBytes - CalcLen(en.value) + CalcLen(value)
		en.value = value
		return
	}
	en := &entry{key, value}
	e := f.ll.PushBack(en)
	f.cache[key] = e
	f.usedBytes += en.Len()
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest()
	}
}
func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}

	return nil
}
func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}
func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Front())
}
func (f *fifo) Len() int {
	return f.ll.Len()
}
func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.usedBytes -= en.Len()
	delete(f.cache, en.key)

	if f.onEvicted != nil {
		f.onEvicted(en.key, en.value)
	}
}
