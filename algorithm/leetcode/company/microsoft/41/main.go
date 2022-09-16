package main

import (
	"fmt"
	"math"
)

/**
思路：一旦数组中出现小于0或者大于n的数，就意味着数组中一定有小于等于n的正数。
*/

func main() {
	var nums = []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	var n = len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] >= math.MaxInt32 || nums[i] <= -math.MaxInt32 {
			continue
		} else {
			index := nums[i]
			if index < 0 {
				index = -index
			}
			index--
			if nums[index] > 0 {
				nums[index] = -nums[index]
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
