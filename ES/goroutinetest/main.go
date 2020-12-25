package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		ch <- i
		fmt.Println("produce:" + strconv.Itoa(i))
		mu.Unlock()
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		v := <-ch
		fmt.Println("consumer:" + strconv.Itoa(v))
		mu.Unlock()
	}
}
func main() {
	ch := make(chan int, 5)
	go produce(ch)
	go consumer(ch)
	time.Sleep(10 * time.Second)
}
