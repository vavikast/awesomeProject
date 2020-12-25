package main

import (
	"fmt"
)

func main() {
	fmt.Println(Fi(6))
}

func Fi(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return Fi(n-1) + Fi(n-2)
}
