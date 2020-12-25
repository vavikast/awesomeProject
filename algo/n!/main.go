package main

import "fmt"

func main() {
	fmt.Println(fn(5))
}

func fn(n int) int {
	if n == 1 {
		return 1
	}
	return n * fn(n-1)

}
