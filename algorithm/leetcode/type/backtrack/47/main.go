package main

import "sort"

var myNums []int
var res [][]int
var used []bool

func permuteUnique(nums []int) [][]int {
	res = nil
	myNums = nums
	used = make([]bool, 21) // [-10, 10]
	sort.Ints(myNums)
	dfs([]int{})
	return res
}

func dfs(b []int) {
	if len(b) == len(myNums) {
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
		return
	}
	for i := 0; i < len(myNums); i++ {
		// 如果前一个数与自己相等，而且前一个数字没有被放进b中，说明这次到我这里了。
		if i > 0 && myNums[i-1] == myNums[i] && used[i-1+10] == false {
			continue
		}
		if used[i+10] {
			continue
		}
		used[i+10] = true
		b = append(b, myNums[i])
		dfs(b)
		used[i+10] = false
		b = b[:len(b)-1]
	}
}
