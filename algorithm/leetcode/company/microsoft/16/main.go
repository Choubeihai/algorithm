package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res = math.MaxInt32
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var left = i + 1
		var right = n - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == target {
				return target
			} else if nums[i]+nums[left]+nums[right] > target {
				if math.Abs(float64(res-target)) > math.Abs(float64(nums[i]+nums[left]+nums[right]-target)) {
					res = nums[i] + nums[left] + nums[right]
				}
				right--
			} else {
				if math.Abs(float64(res-target)) > math.Abs(float64(nums[i]+nums[left]+nums[right]-target)) {
					res = nums[i] + nums[left] + nums[right]
				}
				left++
			}
		}
	}
	return res
}
