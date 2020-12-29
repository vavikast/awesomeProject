package main

import "math/rand"

const (
	maxLevel int     = 16
	p        float32 = 0.25
)

//现在定义一下跳跃表的节点。Score 存放可排序键，Value 存放键对应的值，forward 存放向前的指针列表，forward 的长度是插入元素随机生成的长度。
type Element struct {
	Score   float64
	Value   interface{}
	forward []*Element
}

func newElement(score float64, value interface{}, level int) *Element {
	return &Element{
		Score:   score,
		Value:   value,
		forward: make([]*Element, level),
	}
}

//再来定义跳跃表。header 是一个哑节点，作用和单链表的哑节点一样，方便操作，len 记录当前元素个数，level 记录当前所有元素最大 level。
type SkipList struct {
	header *Element
	len    int
	level  int
}

func newSkipList() *SkipList {
	header := newElement(0, nil, maxLevel)
	return &SkipList{header: header}
}

//需要一个生成随机 level 的函数：
func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Insert(score float64, value interface{}) *Element {
	update := make([]*Element, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	level := randomLevel()
	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.header
		sl.level = level
	}
	e := newElement(score, value, level)
	for i := 0; i < level; i++ {
		e.forward[i] = update[i].forward[i]
		update[i].forward[i] = e
	}
	sl.len++
	return e
}
