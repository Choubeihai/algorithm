package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	sort.Ints(nums)
	fmt.Println(nums)
	res := threeSum(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		} else {
			return nil
		}
	}
	var res [][]int
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
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
