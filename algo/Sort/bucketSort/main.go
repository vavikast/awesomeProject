package main

import "fmt"

//桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。 //
//9.1 算法描述
//设置一个定量的数组当作空桶；
//遍历输入数据，并且把数据一个一个放到对应的桶里去；
//对每个不是空的桶进行排序；
//从不是空的桶里把排好序的数据拼接起来。
//参考: https://www.cnblogs.com/onepixel/p/7674659.html
//参考: https://blog.csdn.net/liaoshengshi/article/details/47320023

//1. 设置固定数量的空桶。
//2. 把数据放到对应的桶中。
//3. 对每个不为空的桶中数据进行排序。()
//4. 拼接从不为空的桶中数据，得到结果。

func bucketSort(arr []int, bucketSize int) []int {
	//找到最大和最小值
	temp := make([][]int, 5)
	length := len(arr)
	newArr := make([]int, 0)
	max := arr[0]
	min := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] >= max {
			max = arr[i]
		}
		if arr[i] <= min {
			min = arr[i]
		}
	}
	//设置桶数量比如是5,则每个桶的范围是bucksize。所以说如果桶的数量需要参考数组长度

	//这里加+1是，这个是整数除。为了避免越界要+1  比如10/3 ,应该是 3.3。
	bucketCount := (max-min)/bucketSize + 1
	fmt.Println(bucketCount)
	for i := 0; i < length; i++ {
		x := (arr[i] - min) / bucketCount
		temp[x] = append(temp[x], arr[i])
	}
	for j := 0; j < 5; j++ {
		insertionSort(temp[j])
		for len(temp[j]) > 0 {
			newArr = append(newArr, temp[j][0])
			temp[j] = temp[j][1:]
		}
	}
	return newArr

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
	fmt.Println(bucketSort(arr, 5))
}
