package main

import "fmt"

//计数排序不是基于比较的排序算法，其核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。 作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
//8.1 算法描述
//找出待排序的数组中最大和最小的元素；
//统计数组中每个值为i的元素出现的次数，存入数组C的第i项；
//对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）；
//反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html

func countigSort(arr []int) []int {
	length := len(arr)
	newarr := make([]int, 0)
	temp := make([]int, length)
	max := arr[0]
	min := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] >= max {
			max = arr[i]
		}
		if arr[i] <= min {
			min = arr[i]
		}
		x := arr[i]
		temp[x]++
	}

	for j := min; j <= max; j++ {
		for temp[j] > 0 {
			newarr = append(newarr, j)
			temp[j]--
		}
	}
	return newarr

}

func main() {
	arr := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11, 12, 11, 11, 12}
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Println(countigSort(arr))
}
