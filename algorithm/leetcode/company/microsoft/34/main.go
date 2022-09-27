package main

import "fmt"

func main() {
	var nums = []int{5, 7, 7, 8, 8, 10}
	var target = 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	first := findFirst(nums, 0, n-1, target)
	last := findLast(nums, 0, n-1, target)
	return []int{first, last}
}

func findFirst(nums []int, left, right int, target int) int {
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}
	mid := (left + right) / 2

	if nums[mid] >= target {
		return findFirst(nums, left, mid, target)
	} else {
		return findFirst(nums, mid+1, right, target)
	}
}

func findLast(nums []int, left, right int, target int) int {
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}
	mid := (left + right + 1) / 2
	if nums[mid] <= target {
		return findLast(nums, mid, right, target)
	} else {
		return findLast(nums, left, mid-1, target)
	}
}
