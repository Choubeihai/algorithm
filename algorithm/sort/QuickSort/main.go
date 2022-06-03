package main

/**
快排序 非稳定排序
*/

import "fmt"

func main() {
	var arr = []int{9, 8, 10, 1, 3, 2, 1, 3}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left int, right int) {
	if left < right {
		partition := partition(arr, left, right)
		quickSort(arr, left, partition-1)
		quickSort(arr, partition+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	i := left
	j := right
	pivot := arr[left]
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = pivot
	return i
}
