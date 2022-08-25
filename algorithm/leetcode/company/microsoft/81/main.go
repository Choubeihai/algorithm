package main

import (
	"fmt"
)

func main() {
	nums2 := []int{2, 2, 0, 0, 1}
	target2 := 1
	//nums2 := []int{1, 0, 1, 1, 1}
	//target2 := 0
	fmt.Println(search(nums2, target2))
}

// 和leetcode 33 题很像
func search(nums []int, target int) bool {
	var n = len(nums)
	return find(nums, 0, n-1, target)
}

func find(nums []int, left, right int, target int) bool {
	if left == right {
		if nums[left] == target {
			return true
		} else {
			return false
		}
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
		return find(nums, left+1, right, target)
	}
	if nums[left] <= nums[mid] { // 左侧有序，一定要加=，因为left可能等于mid
		if target >= nums[left] && target <= nums[mid] {
			return find(nums, left, mid, target)
		} else {
			return find(nums, mid+1, right, target)
		}
	} else { // 右侧有序
		if target >= nums[mid+1] && target <= nums[right] {
			return find(nums, mid+1, right, target)
		} else {
			return find(nums, left, mid, target)
		}
	}
}
