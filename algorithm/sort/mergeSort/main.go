package main

import "fmt"

/**
归并排序 稳定排序
*/

func main() {
	var arr = []int{2, 1, 3, 1, 3, 5, 9}
	fmt.Println(arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		// mid := left + (right-left)/2 // 应付大数，因为 left + right 容易溢出
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, right, mid)
	}
}

func merge(arr []int, left int, right int, mid int) {
	arr1 := make([]int, mid-left+1)
	arr2 := make([]int, right-mid)
	copy(arr1, arr[left:mid+1])
	copy(arr2, arr[mid+1:])
	var i, j int
	var index = left
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			arr[index] = arr1[i]
			i++
		} else {
			arr[index] = arr2[j]
			j++
		}
		index++
	}
	for i < len(arr1) {
		arr[index] = arr1[i]
		i++
		index++
	}
	for j < len(arr2) {
		arr[index] = arr2[j]
		j++
		index++
	}
}
