package main

import "fmt"

func main() {
	var arr = []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearchFirstIndex(arr, 0, len(arr)-1, 7))
}

// 寻找元素出现的第一个位置
func binarySearchFirstIndex(arr []int, left, right int, target int) int {
	if left > right {
		return -1
	}
	if arr[left] == target {
		return left
	}

	mid := (left + right) / 2
	if arr[mid] > target {
		return binarySearchFirstIndex(arr, left, mid-1, target)
	} else if arr[mid] == target {
		return binarySearchFirstIndex(arr, left, mid, target)
	} else {
		return binarySearchFirstIndex(arr, mid+1, right, target)
	}
}

// 寻找元素出现的最后一个位置
func binarySearchLastIndex(arr []int, left, right int, target int) int {
	if left > right {
		return -1
	}
	if arr[left] == target {
		return left
	}

	mid := (left + right + 1) / 2
	if arr[mid] > target {
		return binarySearchLastIndex(arr, left, mid-1, target)
	} else if arr[mid] == target {
		return binarySearchLastIndex(arr, mid, right, target)
	} else {
		return binarySearchLastIndex(arr, mid+1, right, target)
	}
}
