package main

import (
	"math"
)

const (
	MAX_LEVEL = 16
)

type skipListNode struct {
	//跳表保存的值
	v interface{}
	//用于排序的分值
	score int
	//层高
	level int
	//每层前进的指针
	forwards []*skipListNode
}

//新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
}

//跳表结构体
type SkipList struct {
	//跳表头节点
	head *skipListNode
	//调表当前层数
	level int
	//调表长度
	length int
}

func NewSkipList() *SkipList {
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}

}

//获取调表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//获取跳表级数
func (sl *SkipList) Level() int {
	return sl.level
}

//插入节点到跳表

func (sl *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找插入位置
	cur := sl.head
	//记录每层的路径
	update := [MAX_LEVEL]*skipListNode{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}
}

func main() {

}
