package main

import "fmt"

func main() {
	var nums = []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}

//len(nums)=n： 一旦数组中出现小于0或者大于n的数，就意味着数组中一定有小于等于n的正数确实。
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if index < len(nums) {
			nums[index] = -abs(nums[index])
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
