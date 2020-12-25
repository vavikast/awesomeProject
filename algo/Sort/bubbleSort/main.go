package main

//冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
//1.1 算法描述
//比较相邻的元素。如果第一个比第二个大，就交换它们两个；
//对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
//针对所有的元素重复以上的步骤，除了最后一个；
//重复步骤1~3，直到排序完成。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html
import "fmt"

func bulleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func bulleSort1(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func main() {
	a := []int{1, 3, 5, 4, 2}
	fmt.Println(bulleSort(a))
	b := []int{1, 3, 5, 4, 2}
	bulleSort1(b)
	fmt.Println(b)

}
