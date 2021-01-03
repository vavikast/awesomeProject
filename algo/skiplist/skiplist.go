package main

import (
	"fmt"
	"math/rand"
)

//参考：https://juejin.cn/post/6896645943572037640
//数据结构真是一个神奇的东西，刚学的时候，因学不会痛苦的死去活来，学会了又为其精妙的设计惊叹狂喜。

//这个跳表只是个思路。有几个点还是不需要补充和完善
// Interface 空接口  fakeNode 伪节点的问题
// Less Equal没有实现。

//定义最大层数限制
const MAX_LEVEL = 16

//定义层级因子
const LEVEL_FACTOR = 0.5

//定义几种错误代码
const (
	OK = iota + 1
	DUPLICATED
	NOT_EXIST
	NOT_INIT
)

//定义一个描述数据的接口
type Interface interface {
	// less than p
	Less(p Interface) bool
	// equal with p
	Equal(p Interface) bool
}

// 定义了一个伪节点，用于当做链表中的哨兵,主要是实现接口
type fakeNode struct {
}

func (f *fakeNode) Less(p Interface) bool {
	return false
}

func (f *fakeNode) Equal(p Interface) bool {
	return false
}

//链表节点
type node struct {
	data     Interface //实际的数据
	forwards []*node   //索引存储的位置 //这个切片对应的是Level. 如果在第二层的第一个节点:head.forwards[2](存放的是地址，指向下一个节点的地址).第二层的第二个节点head.forwards[2].forwards[2]
	level    int       //节点的层级
}

//链表定义
type SkipList struct {
	head   *node  //链表头节点
	length uint32 //链表长度
	level  int    //当前跳表最高级别

}

//初始化一个节点
func newNode(p Interface, l int) *node {
	return &node{
		data:     p,
		forwards: make([]*node, l, l),
		level:    l,
	}

}

//c初始化一个跳表，head默认拥有最高级别
func NewSkipList() *SkipList {
	return &SkipList{newNode(&fakeNode{}, MAX_LEVEL), 0, 1}
}

//下面主要试下一下三个方法 Add，Delete和Search。

//跳表添加元素
//这里随机产生一个层级
//在LEVEL_FACTOR是0.5的情况下
// 1 级的概率是 50%
// 2 级的概率是 25%
// 3 级的概率是 12.5%, 以此类推

func (sl *SkipList) randomLevel() int {
	l := 1
	for rand.Float64() < LEVEL_FACTOR && l < MAX_LEVEL {
		l++
	}

	//如果层级比当前层级高2级或2级以上，按照高一级处理，避免浪费。（这个实在是高）
	if sl.level+1 < l {
		return sl.level + 1 //#跳表的第一个元素是否一次加到位。
	}
	return l
}

//增加一级到跳表
func (sl *SkipList) Add(p Interface) int {
	//如果head为空，则说明没有初始化
	if sl.head == nil {
		return NOT_INIT
	}

	//找到应该插入的位置，从头节点开始
	cur := sl.head
	//查找路径记录 。 比如跳表总共有三级索引，那么的话 如果数据在二级索引开始找到，这个时候需要下探一层继续找，最后定位到具体的位置。所以必须要记录2级下探索引和1级索引所在节点位置。
	update := [MAX_LEVEL]*node{}

	//从最高级开始查找.head始终是指向链顶节点的。
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		// 从head当前层向后查找。。
		for nil != cur.forwards[i] {
			//如果相等，则不允许插入
			if cur.forwards[i].data.Equal(p) {
				return DUPLICATED
			}
			//若插入元素小于后一个元素的值。则意味着要向下寻找。下探到下一层
			if !cur.forwards[i].data.Less(p) {
				//记录当前层的地址node地址
				update[i] = cur
				break
			}
			//若当前元素小于要插入的元素，下一个元素也小于于要插入的元素，则意味着要同层向后移动。
			cur = cur.forwards[i]
		}
		//如果这层遍历到结束，还没有找到对应的位置
		//那么就将最后的元素作为当前层级路径
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//生成当前节点层级
	l := sl.randomLevel()
	//初始化节点
	n := newNode(p, l)

	//从最底层开始添加节点和索引。
	for i := 0; i < n.level; i++ {
		n.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = n
	}

	//更新调表的索引
	if n.level > sl.level {
		sl.level = n.level
	}
	return OK

}

//删除跳表元素
func (sl *SkipList) Delete(p Interface) int {
	//查找元素
	//与添加不同的时，不需要考虑相等的情况
	//如果某一个层级没有的话，不需要记录层级最后的节点。
	cur := sl.head
	update := [MAX_LEVEL]*node{}
	for i := sl.level; i >= 0; i-- {
		update[i] = cur
		for nil != cur.forwards[i] {
			if !cur.forwards[i].data.Less(p) {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	//找到要删除的节点
	cur = update[0].forwards[0]
	//没有这个元素
	if cur == nil {
		return NOT_EXIST
	}
	//从节点所在最高级依次向下删除
	for i := cur.level - 1; i >= 0; i-- {
		//如果当前的节点是某一个层级的最后一个元素
		//那么就就降低跳表的层级
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}
		//删除元素
		if nil != update[i].forwards[i] {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}
	sl.length--
	return OK
}

//查找跳表的元素
func (sl *SkipList) Search(p Interface) *node {
	cur := sl.head
	update := [MAX_LEVEL]*node{}
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].data.Equal(p) {
				return cur
			}
			if !cur.forwards[i].data.Less(p) {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}

		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}
	return nil
}

// 按层级打印跳表
func (sl *SkipList) Print() {
	cur := sl.head
	for i := sl.level; i >= 0; i-- {
		fmt.Printf("[level = %d] ", i)
		for nil != cur {
			fmt.Printf("%+v   ", cur.data)
			cur = cur.forwards[i]
		}
		fmt.Println("")
		cur = sl.head
	}
}

// 获取长度
func (sl *SkipList) Length() uint32 {
	return sl.length
}

// 获取高度
func (sl *SkipList) Level() int {
	return sl.level
}
