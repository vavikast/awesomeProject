package main

//归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
//5.1 算法描述
//把长度为n的输入序列分成两个长度为n/2的子序列；
//对这两个子序列分别采用归并排序；
//将两个排序好的子序列合并成一个最终的排序序列。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html
//参考: https://blog.csdn.net/k_koris/article/details/80508543
import "fmt"

//先分后合，就是先mergeSort到底（递归终止条件），再一步步合并回来merge回来。
func mergeSort(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	i := len / 2
	leftArr := arr[:i]
	rightArr := arr[i:]
	return merge(mergeSort(leftArr), mergeSort(rightArr))

}
func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	//承接上文。只剩下下面两种情况，根据二分法的偏差，left存在,还剩一个，right不存在；left不存在，right存在，还剩一个。 下面是默认省略了len(left|right) == 0
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func main() {
	a := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11, 12}
	fmt.Println(mergeSort(a))

}
