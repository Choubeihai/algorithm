package main

//len(nums)=n： 一旦数组中出现小于0或者大于n的数，就意味着数组中一定有小于等于n的正数确实。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		index := nums[i] - 1
		if index < n {
			nums[index] = -abs(nums[index])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
