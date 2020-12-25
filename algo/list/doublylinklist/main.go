package main

import "fmt"

//双链表

type Node struct {
	property     interface{}
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

type LinkedLister interface {
	AddToHead(interface{})
	NodeWithValue(interface{}) *Node
	AddAfter(interface{}, interface{})
	LastNode() *Node
	AddToEnd(interface{})
	GetNextValues(interface{}) interface{}
	GetPreviousValues(interface{}) interface{}
	IterateList()
	Clear()
	Len() int
}

//添加到头节点

// 添加头节点就是理解的难点
// 初始化时 链表中没有node。这个链接指向了nil。
// 空node的pre和next都是nil
//如果双向链表为空。headnode直接指向node1。将node替换成headnode就行。因为只有一个节点，而且node1初始化时pre和next都是nil，所以直接替换headnode即可。
//第一步初始化一个node。 存储区域赋值，pre和next均为nil。

//如果双向链表不空。将node2加入链表。就相当于将headnode的指向从node1指向node2.此时
//第一步初始化一个node2。存储区域赋值，此时的node2的pre和next均为nil。
//因为这个node将来是头节点，所以这个节点的pre将来也会是nil，我们不做调整。
//根据链表插入的原则，将现在headnode的pre指向node（原来为nil） 即为 l.headNode.previousNode = node
//设置node的后续节点为现在的headNode. 即为 node.nextNode = l.headNode
//最后将headnode的指向变为node。

func (l *LinkedList) AddToHead(property interface{}) {
	//声明并进行了初始化。可以使用new(&Node{}代替)
	var node = &Node{}
	node.property = property
	//如果当前不是空节点
	if l.headNode != nil {
		l.headNode.previousNode = node
		node.nextNode = l.headNode

	}
	//如果当前是空节点
	l.headNode = node

}

// 查找某值节点
func (l *LinkedList) NodeWithValue(property interface{}) *Node {
	var node *Node
	var withnode *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			withnode = node
			break
		}
	}
	return withnode
}

//某个值后面添加节点
func (l *LinkedList) AddAfter(nodeproperty interface{}, property interface{}) {
	//声明并进行了初始化。可以使用new(&Node{}代替)
	var node = &Node{}
	var withnode *Node
	node.property = property
	withnode = l.NodeWithValue(nodeproperty)
	if withnode != nil {
		withnode.nextNode.previousNode = node
		node.nextNode = withnode.nextNode
		node.previousNode = withnode
		withnode.nextNode = node
	}

}

//寻找最后一个节点
func (l *LinkedList) LastNode() *Node {
	//声明并进行了初始化。可以使用new(&Node{}代替)
	var node *Node
	var lastnode *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastnode = node
		}
	}
	return lastnode
}

//将节点添加到最后
func (l *LinkedList) AddToEnd(property interface{}) {
	//声明并进行了初始化。可以使用new(&Node{}代替)
	var node = &Node{}
	var lastnode *Node
	node.property = property
	lastnode = l.LastNode()
	if lastnode != nil {
		lastnode.nextNode = node
		node.previousNode = lastnode
	}

}
func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}

}

//清空节点
func (l *LinkedList) Clear() {
	l.headNode = nil
}

//获取节点长度
func (l *LinkedList) Len() int {
	var total int
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		total += 1
	}
	return total
}

//获取后一个节点的值
func (l *LinkedList) GetNextValues(property interface{}) interface{} {
	var withnode *Node
	withnode = l.NodeWithValue(property)
	if withnode.nextNode != nil {
		return withnode.nextNode.property
	}
	return nil
}

//获取前一个节点的值
func (l *LinkedList) GetPreviousValues(property interface{}) interface{} {
	var withnode *Node
	withnode = l.NodeWithValue(property)
	if withnode.previousNode != nil {
		return withnode.previousNode.property
	}
	return nil
}
func main() {
	//只是为了调用接口
	var l LinkedLister
	l = &LinkedList{}
	fmt.Println(l)
	//初始化headnode
	linkedList := LinkedList{}
	//在开头添加节点
	linkedList.AddToHead(8)
	//fmt.Println(linkedList.headNode.property)
	linkedList.AddToHead(4)
	//fmt.Println(linkedList.headNode.property)
	//fmt.Println(linkedList.headNode.nextNode.property)
	//linkedList.IterateList()
	//在末尾添加节点
	linkedList.AddToEnd(11)
	linkedList.AddToEnd(15)
	//在11后面添加添加节点的值为12
	linkedList.AddAfter(11, 12)
	//查看节点

	linkedList.IterateList()
	fmt.Println("The Length is :")
	fmt.Println(linkedList.Len())
	fmt.Println("The number next  is :")
	fmt.Println(linkedList.GetNextValues(11))
	fmt.Println("The number previous  is :")
	fmt.Println(linkedList.GetPreviousValues(11))
}
