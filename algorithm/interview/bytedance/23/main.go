package main

import "fmt"

/**
两个有序数组找第k大，要求O(logk)时间复杂度
*/

func main() {
	var arr1 = []int{1, 2, 3, 4}
	arr2 := []int{3, 5, 7, 8}
	fmt.Println(findKth(arr1, arr2, 5))
}

func findKth(arr1 []int, arr2 []int, k int) int {
	if len(arr1) == 0 {
		return arr2[k-1]
	}
	if len(arr2) == 0 {
		return arr1[k-1]
	}

	if k == 1 {
		return min(arr1[0], arr2[0])
	}
	k1 := min(k/2, len(arr1))
	k2 := min(k-k1, len(arr2))
	if arr1[k1-1] > arr2[k2-1] {
		return findKth(arr1, arr2[k2:], k-k2)
	} else {
		return findKth(arr1[k1:], arr2, k-k1)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
