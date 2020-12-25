package main

import "fmt"

//插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
//3.1 算法描述
//一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：
//从第一个元素开始，该元素可以认为已经被排序；
//取出下一个元素，在已经排序的元素序列中从后向前扫描；
//如果该元素（已排序）大于新元素，将该元素移到下一位置；
//重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
//将新元素插入到该位置后；
//重复步骤2~5。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html
func insertionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
func insertionSort1(arr []int) {
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
	a := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(insertionSort(a))
	b := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	insertionSort1(b)
	fmt.Println(b)

}
