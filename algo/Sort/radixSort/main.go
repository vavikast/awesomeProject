package main

import (
	"fmt"
	"math"
)

//基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。
//10.1 算法描述
//取得数组中的最大数，并取得位数；
//arr为原始数组，从最低位开始取每个位组成radix数组；
//对radix进行计数排序（利用计数排序适用于小范围数的特点）；
//参考: https://www.cnblogs.com/onepixel/p/7674659.html

func radixSort(arr []int) []int {
	//找到最大值和位数
	length := len(arr)
	max := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] >= max {
			max = arr[i]
		}
	}
	n := 0
	if max/10 > 0 {
		n++
	}

	for i := 0; i <= n; i++ {
		newArr := make([]int, 0)
		temp := make([][]int, length)
		x := int(math.Pow10(i))
		for j := 0; j < length; j++ {

			num := arr[j] / x % 10
			temp[num] = append(temp[num], arr[j])

		}

		for y := 0; y < 10; y++ {
			insertionSort(temp[y])
			for len(temp[y]) > 0 {
				newArr = append(newArr, temp[y][0])
				temp[y] = temp[y][1:]
			}
		}
		arr = newArr

	}
	//

	return arr
}

func insertionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

}
func main() {
	arr := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11, 12, 11, 11, 12}
	fmt.Println(len(arr))
	fmt.Println(radixSort(arr))
}
