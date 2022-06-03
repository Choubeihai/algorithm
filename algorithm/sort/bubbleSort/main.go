package main

/**
冒牌排序 稳定排序
*/

import "fmt"

func main() {
	var arr = []int{9, 8, 10, 1, 3, 2, 1, 3}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
