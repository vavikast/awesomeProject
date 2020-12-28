package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	array []interface{}
	size  int
	lock  sync.Mutex
}

func NewArrayStack() *ArrayStack {
	NewArrayStack := &ArrayStack{
		array: make([]interface{}, 0),
		size:  0,
		lock:  sync.Mutex{},
	}
	return NewArrayStack
}

func (stack *ArrayStack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.array = append(stack.array, value)
	stack.size++
}
func (stack *ArrayStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("empty")
	}
	v := stack.array[stack.size-1]
	stack.array = stack.array[0 : stack.size-1]
	stack.size--
	return v
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() interface{} {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
func main() {
	arrayStack := NewArrayStack()
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
