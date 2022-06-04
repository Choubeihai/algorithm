package main

import (
	"fmt"
)

/**
将数组一分为二，其中一定有一个是有序的，另一个可能有序，也能是部分有序。
此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环.
*/

func main() {
	var nums = []int{1}
	target := 0
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[left] == target {
		return left
	}
	if nums[mid] == target {
		return mid
	}
	if nums[right] == target {
		return right
	}

	if nums[left] < nums[mid] { // 左边有序
		if target > nums[left] && target < nums[mid] {
			return binarySearch(nums, left+1, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right-1, target)
		}
	} else {
		if target > nums[mid] && target < nums[right] {
			return binarySearch(nums, mid+1, right-1, target)
		} else {
			return binarySearch(nums, left+1, mid-1, target)
		}
	}
}
