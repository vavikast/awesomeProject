package main

import (
	"fmt"
)

//快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
//6.1 算法描述
//快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：
//从数列中挑出一个元素，称为 “基准”（pivot）；
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html
//参考： https://wiki.jikexueyuan.com/project/easy-learn-algorithm/fast-sort.html

func quickSort(arr []int) []int {
	return subQuickSort(arr, 0, len(arr)-1)
}

//也是分治方法，先分后合，意味着分有个终止条件，才能合的上。
func subQuickSort(arr []int, low, high int) []int {
	if low < high {
		//分区点本身已经是最优数据了，所以不需要排序
		pivot := partition(arr, low, high)
		subQuickSort(arr, low, pivot-1)
		subQuickSort(arr, pivot+1, high)
	}
	//
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := low
	for low < high {
		if arr[high] >= arr[pivot] && high > 0 {
			high--
		}
		if arr[low] < arr[pivot] {
			low++
		}
		if low < high {
			arr[high], arr[low] = arr[low], arr[high]
		}
	}
	arr[high], arr[pivot] = arr[pivot], arr[high]
	pivot = high
	return pivot
}

func main() {
	a := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11, 12, 11, 11, 12}
	fmt.Println(quickSort(a))

}
