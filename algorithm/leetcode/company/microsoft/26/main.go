package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	var slow = 1
	var fast = 1
	for fast = 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
