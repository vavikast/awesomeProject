package main

import "fmt"

//1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序。
//4.1 算法描述
//先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：
//选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
//按增量序列个数k，对序列进行k 趟排序；
//每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html
func shellSort(arr []int) []int {
	len := len(arr)
	gap := len / 2
	//每次选择gap的一半
	for i := gap; i > 0; i = i / 2 {
		for j := 0; j <= i; j++ {
			for x := j; x < len-i; x = x + i {
				if arr[x] > arr[x+i] {
					arr[x], arr[x+i] = arr[x+i], arr[x]
				}
			}
		}
	}
	return arr
}

func shellSort1(arr []int) {
	len := len(arr)
	gap := len / 2
	//每次选择gap的一半
	for i := gap; i > 0; i = i / 2 {
		for j := 0; j <= i; j++ {
			for x := j; x < len-i; x = x + i {
				if arr[x] > arr[x+i] {
					arr[x], arr[x+i] = arr[x+i], arr[x]
				}
			}
		}
	}

}

func main() {
	a := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11, 12}
	fmt.Println(shellSort(a))
	b := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11}
	shellSort1(b)
	fmt.Println(b)

}
