package main

import "fmt"

func main() {
	var arr = []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearch(arr, 0, len(arr)-1, 7))
}

func binarySearch(arr []int, left, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[left] == target {
		return left
	}
	if arr[mid] == target {
		return mid
	}
	if arr[right] == target {
		return right
	}

	if arr[mid] > target {
		return binarySearch(arr, left+1, mid-1, target)
	} else {
		return binarySearch(arr, mid+1, right-1, target)
	}

}
