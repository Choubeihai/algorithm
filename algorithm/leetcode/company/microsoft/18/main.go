package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{1, 0, -1, 0, -2, 2}
	var target = 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			var left = j + 1
			var right = n - 1
			for left < right {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right {
						if nums[left] == nums[left-1] {
							left++
						} else {
							break
						}
					}
					right--
					for left < right {
						if nums[right] == nums[right+1] {
							right--
						} else {
							break
						}
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
