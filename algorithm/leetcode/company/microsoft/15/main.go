package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		var left = i + 1
		var right = n - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
