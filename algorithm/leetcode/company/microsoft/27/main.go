package main

func removeElement(nums []int, val int) int {
	var slow = 0
	var fast = 0
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
