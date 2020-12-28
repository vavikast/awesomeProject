package main

func BinarySearch(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] >= v {
			high = mid - 1
		} else {
			low = mid + 1
		}
		if low < n-1 && a[low] == v {
			return v
		}
	}

	return -1
}
