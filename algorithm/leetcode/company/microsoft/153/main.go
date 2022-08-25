package main

func findMin(nums []int) int {
	var n = len(nums)
	return find(nums, 0, n-1)
}

func find(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	if nums[left] <= nums[mid] { // 左侧有序 [left, mid]
		a := find(nums, mid+1, right) // 接着寻找右侧 [mid+1, right]
		return min(a, nums[left])
	} else { // 右侧有序 [mid+1, right]
		a := find(nums, left, mid) // 接着寻找左侧[left, mid]
		return min(a, nums[mid+1])
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
