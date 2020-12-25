package main

import (
	"fmt"
)

var total int

func main() {
	n := 3
	MoveTover(n, "A", "B", "C")
	fmt.Println(total)
}

//如果有n>=3 借助B到C
//递归总有终止条件的，只要找到终止条件即可。
//前n-1个.要先A挪到B。A->B 递归
//A->C 递归
//A->B一步
//C->B 递归
//然后将A第n个挪到C。 A->C  一步
//再把B中第n-1个挪到C。B->C  递归
//B->A 递归
//B->C一步
//A->C 递归

func MoveTover(n int, A, B, C string) {
	if n == 1 {
		total = total + 1
		fmt.Println(A + "->" + C)
		return
	}
	MoveTover(n-1, A, C, B)
	fmt.Println(A + "->" + C)
	total = total + 1
	MoveTover(n-1, B, A, C)
}
