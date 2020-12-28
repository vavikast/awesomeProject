package main

import (
	"fmt"
	"sync"
)

type ArrayQueue struct {
	array []interface{}
	size  int
	lock  sync.Mutex
}

func NewArrayQueue() *ArrayQueue {
	ArrayQueue := &ArrayQueue{
		array: nil,
		size:  0,
		lock:  sync.Mutex{},
	}
	ArrayQueue.array = make([]interface{}, 0)
	return ArrayQueue
}

func (queue *ArrayQueue) Add(value interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.array = append(queue.array, value)
	queue.size++
}
func (queue *ArrayQueue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		panic("empty")
	}
	v := queue.array[0]
	queue.array = queue.array[1:]
	queue.size--
	return v
}
func main() {
	q := NewArrayQueue()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)
	q.Remove()
	fmt.Println(q.array)

}
