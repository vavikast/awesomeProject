package main

import "fmt"

//7、堆排序（Heap Sort）
//堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
//7.1 算法描述
//将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
//将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
//由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
//参考：https://www.cnblogs.com/onepixel/p/7674659.html

//小顶堆
func HeapSort(arr []int) {
	//求数组长度
	//根据堆的规律，假设子节点的规律，假设子节点的坐标为i
	//左子节点坐标为2*i+1,右子节点坐标为2*i+2
	//父节点的坐标为（i-1）/2.   此处可以计算无论最后一位数字在做左子节点，还是右子节点。父节点的坐标一定是（i-1）/2。 golang中/取整
	//假设切片长度是len（arr），那么最后一位的坐标序号为len(arr)-1,可计算出父节点的位置为（len(arr)-1）/2
	length := len(arr)
	last_node := length - 1
	//建立最大堆，最大堆的概念就是父节点总是比子节点数字大。arr[0]最大
	buildHeap(arr)
	//比如 312 -> 21  3-->1  2   3=123
	for i := last_node; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr[:i], 0)
	}
}

func buildHeap(arr []int) {
	length := len(arr)
	last_node := length - 1
	parent := (last_node - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(arr, i)

	}
}
func heapify(arr []int, i int) {
	length := len(arr)
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left < length && arr[left] > arr[max] {
		max = left
	}
	if right < length && arr[right] > arr[max] {
		max = right
	}
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, max)
		//此处是向下递归，目的就是建立最大堆，因为比较的过程并不能保证最大堆,只是让他们换了位置。

	}
}
func main() {
	a := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 11, 12, 11, 11, 12}
	HeapSort(a)
	fmt.Println(a)
}
