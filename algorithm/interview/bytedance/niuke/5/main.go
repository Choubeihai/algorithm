package main

import (
	"fmt"
	"math"
	"sort"
)

/**
正整数数组nums，求和为target的最短子数组
*/

func main() {
	nums := []int{1, 2, 3, -1, 3}
	target := 3
	fmt.Println(findArr(nums, target))
}

var minLength int
var res []int

func findArr(nums []int, target int) []int {
	sort.Ints(nums)
	minLength = math.MaxInt32
	res = nil
	dfs(nums, 0, []int{}, target)
	return res
}

func dfs(nums []int, index int, arr []int, target int) {
	if index >= len(nums) {
		return
	}
	if target < 0 {
		return
	}
	if target == 0 {
		if len(arr) < minLength {
			minLength = len(arr)
			c := make([]int, len(arr))
			copy(c, arr)
			res = c
		}
		return
	}

	for i := 0; i < len(nums); i++ {
		if i < index {
			continue
		}
		target = target - nums[i]
		arr = append(arr, nums[i])
		dfs(nums, i+1, arr, target)
		target = target + nums[i]
		arr = arr[:len(arr)-1]
	}

}
