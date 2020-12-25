package main

import "fmt"

type Node struct {
	property interface{}
	nextNode *Node
}

//永远锚定了头结点。
type LinkedList struct {
	headNode *Node
}

type LinkedLister interface {
	AddToHead(interface{})
	IterateList()
	LastNodeData() interface{}
	LastNode() *Node
	AddToEnd(interface{})
	NodeWithValue(interface{}) *Node
	AddAftermethod(nodeproerty, proerty interface{})
	Clear()
	Length()
}

//  head  {nil,nil}
//  添加第一个数据则变为 LinkedList{data1,nil}
//  添加第二个数据则变为 {data2,&LinkedList{data1,nil}}
//  添加第三个数据则变为 {data3,&LinkedList{data2,&LinkedList{data1,nil}}}

// 添加头结点
func (linkedList *LinkedList) AddToHead(proerty interface{}) {
	var node = Node{}
	node.property = proerty
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

//迭代器 （遍历节点）
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// 查看最后一个节点的数据
func (linkedList *LinkedList) LastNodeData() interface{} {
	var node *Node
	node = linkedList.headNode
	for {
		node = node.nextNode
		if node.nextNode == nil {
			return node.property
		}
	}

}

// 寻找最后一个节点
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode

}

// 添加尾节点
func (linkedList *LinkedList) AddToEnd(proerty interface{}) {
	var endNode = &Node{}
	var lastNode *Node
	lastNode = linkedList.LastNode()
	endNode.property = proerty
	endNode.nextNode = nil
	if lastNode != nil {
		lastNode.nextNode = endNode
	}
}

// 根据存储值，查找所在节点
func (linkedList *LinkedList) NodeWithValue(property interface{}) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

//插入到特定值后面 .如果要插入链表，必须先将以拆开链表的后端接到新的节点，再把新的节点接入到链接上
func (linkedList *LinkedList) AddAftermethod(nodeproerty, proerty interface{}) {
	var node = &Node{}
	node.property = proerty
	node.nextNode = nil
	var withnode *Node
	withnode = linkedList.NodeWithValue(nodeproerty)
	if withnode != nil {
		node.nextNode = withnode.nextNode
		withnode.nextNode = node
	}

}

//清空链表
func (linkedList *LinkedList) Clear() {
	linkedList.headNode = nil
}

//链表长度
func (linkedList *LinkedList) Length() int {
	var node *Node
	var total int
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		total += 1
	}
	return total
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(5)
	//查看头结点的值
	fmt.Println("查看头节点:", linkedList.headNode.property)

	// 查看迭代器
	linkedList.IterateList()

	//在最后插入一个节点
	linkedList.AddToEnd(4)
	linkedList.AddToEnd("string")
	//查看迭代器
	linkedList.IterateList()
	//查看存值位置.(目的就是验证interface{}作用)
	fmt.Println("查看存储数值", linkedList.NodeWithValue("string").property.(string))
	// 在5后面插入4.
	linkedList.AddAftermethod(5, 4)
	//查看迭代器
	linkedList.IterateList()
	//查看链表长度
	fmt.Println("the linkedlist 长度为: ", linkedList.Length())

	fmt.Println("清空链表")
	linkedList.Clear()
	linkedList.AddToHead(1)
	// 查看迭代器
	linkedList.IterateList()
	//查看链表长度
	fmt.Println("the linkedlist 长度为: ", linkedList.Length())

}
