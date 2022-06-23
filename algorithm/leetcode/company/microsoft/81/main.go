package main

import (
	"fmt"
)

func main() {
	nums2 := []int{2, 2, 2, 0, 0, 1}
	target2 := 0
	//nums2 := []int{1, 0, 1, 1, 1}
	//target2 := 0
	fmt.Println(search(nums2, target2))
}

// 和leetcode 33 题很像
func search(nums []int, target int) bool {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right int, target int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return true
	}
	if nums[left] == target {
		return true
	}
	if nums[right] == target {
		return true
	}
	if nums[left] == nums[mid] {
		return binarySearch(nums, left+1, right, target)
	}

	if nums[left] < nums[mid] {
		if nums[left] <= target && nums[mid] > target {
			return binarySearch(nums, left, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right, target)
		}
	} else {
		if nums[mid] < target && nums[right] >= target {
			return binarySearch(nums, mid+1, right, target)
		} else {
			return binarySearch(nums, left, mid-1, target)
		}
	}
}
