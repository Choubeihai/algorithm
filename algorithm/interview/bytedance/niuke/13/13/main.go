package main

import "fmt"

/**
一个整数数组，长度为n，将其分为m份，使各份的和相等，求m的最大值
    比如{3，2，4，3，6} 可以分成{3，2，4，3，6} m=1;
    {3,6}{2,4,3} m=2
    {3,3}{2,4}{6} m=3 所以m的最大值为3
*/

func main() {
	var arr = []int{3, 2, 4, 3, 6}
	fmt.Println(ArrayPartition(arr))
}

func ArrayPartition(nums []int) int {
	var n = len(nums)
	var m = n
	var sum int
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	for m > 0 {
		if sum%m != 0 {
			m--
		} else {
			subSum := sum / m
			tag := make([]bool, n)
			var i = 0
			for ; i < n; i++ {
				if tag[i] {
					continue
				}
				tag[i] = true
				if !dfs(nums, tag, subSum-nums[i]) {
					break
				}
			}
			if i == n {
				break
			}
			m--
		}
	}
	return m
}

func dfs(nums []int, tag []bool, sum int) bool {
	if sum < 0 {
		return false
	} else if sum == 0 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if tag[i] {
			continue
		} else {
			tag[i] = true
			if dfs(nums, tag, sum-nums[i]) {
				return true
			} else {
				tag[i] = false
				return false
			}
		}
	}
	return false
}
