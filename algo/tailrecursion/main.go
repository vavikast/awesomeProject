package main

import "fmt"

func main() {
	fmt.Println(fn(5, 1))
}

func fn(n int, a int) int {
	if n == 1 {
		return a
	}
	return fn(n-1, a*n)
}
