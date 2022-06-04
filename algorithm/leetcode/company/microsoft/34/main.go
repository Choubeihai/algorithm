package main

import "fmt"

func main() {
	var nums = []int{5, 7, 7, 8, 8, 10}
	var target = 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	firstIndex := binarySearchFirstIndex(nums, left, right, target)
	if firstIndex == -1 {
		return []int{-1, -1}
	}
	lastIndex := binarySearchLastIndex(nums, left, right, target)
	return []int{firstIndex, lastIndex}
}

// 寻找target的第一个位置
func binarySearchFirstIndex(nums []int, left, right int, target int) int {
	if left > right {
		return -1
	}
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}
	mid := (left + right) / 2
	if nums[mid] < target {
		return binarySearchFirstIndex(nums, mid+1, right, target)
	} else if nums[mid] == target {
		return binarySearchFirstIndex(nums, left, mid, target)
	} else {
		return binarySearchFirstIndex(nums, left, mid-1, target)
	}
}

// 寻找target的最后一个位置
func binarySearchLastIndex(nums []int, left, right int, target int) int {
	if left > right {
		return -1
	}
	if left == right {
		return left
	}
	mid := (left + right + 1) / 2
	if nums[mid] < target {
		return binarySearchLastIndex(nums, mid+1, right, target)
	} else if nums[mid] == target {
		return binarySearchLastIndex(nums, mid, right, target)
	} else {
		return binarySearchLastIndex(nums, left, mid-1, target)
	}
}
