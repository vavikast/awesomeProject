package main

import "fmt"

//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。

var total = 0

func main() {
	fmt.Println(FrogJump(3))
}

//小青蛙的最后一步。
// 可以是1步，也可以是2步
//当调n个台阶时。
//前一个台阶可能是n-1, 递归
//前一个台阶可能是n-2, 递归

func FrogJump(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	return FrogJump(n-1) + FrogJump(n-2)
}
