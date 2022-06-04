package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right int, target int) int {
	if left > right {
		return left
	}
	if left == right {
		if nums[left] == target {
			return left
		} else if nums[left] > target {
			return left
		} else {
			return left + 1
		}
	}
	mid := (left + right) / 2
	if nums[mid] > target {
		return binarySearch(nums, left, mid-1, target)
	} else if nums[mid] == target {
		return binarySearch(nums, left, mid, target)
	} else {
		return binarySearch(nums, mid+1, right, target)
	}
}
