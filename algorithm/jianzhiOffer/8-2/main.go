package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	var n = len(nums)
	var i = 0
	var j = 0
	var sum = 0
	var res = math.MaxInt32
	for j < n {
		sum += nums[j]
		for sum >= target {
			res = min(res, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	if res == math.MaxInt32 {
		return 0
	} else {
		return res
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
