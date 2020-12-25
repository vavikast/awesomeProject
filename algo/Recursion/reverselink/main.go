package main

type Node struct {
	Data interface{}
	Next *Node
}

func main() {

}

type Noder interface {
	New() *Node
	Set(x interface{}) *Node
	Get(n int) interface{}
	Len() int
}

func (n *Node) New() *Node {
	node := Node{
		Data: nil,
		Next: nil,
	}
	return &node
}

func (n *Node) Set(x interface{}) *Node {
	if n.Next != nil {
		n = n.Next
	}
	n1 := &Node{
		Data: x,
		Next: nil,
	}
	n.Next = n1
	return n

}

func (n *Node) Get(num int) interface{} {
	panic("implement me")
}

func (n *Node) Len() int {
	panic("implement me")
}
