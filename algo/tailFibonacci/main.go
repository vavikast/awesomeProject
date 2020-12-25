package main

import "fmt"

func main() {
	fmt.Println(Fi(1, 1, 1))
	fmt.Println(Fi(2, 1, 1))
	fmt.Println(Fi(3, 1, 1))
	fmt.Println(Fi(4, 1, 1))
	fmt.Println(Fi(5, 1, 1))
}

func Fi(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return Fi(n-1, a2, a1+a2)
}
