package main

import "sort"

var myNums []int
var visit []bool
var res [][]int

func subsetsWithDup(nums []int) [][]int {
	myNums = nums
	sort.Ints(myNums)
	visit = make([]bool, len(nums))
	res = nil
	dfs([]int{}, 0)
	return res
}

func dfs(b []int, index int) {
	c := make([]int, len(b))
	copy(c, b)
	res = append(res, c)

	if index >= len(myNums) {
		return
	}

	for i := 0; i < len(myNums); i++ {
		if index > i {
			continue
		}
		if i > 0 && myNums[i-1] == myNums[i] && !visit[i-1] {
			continue
		}
		b = append(b, myNums[i])
		visit[i] = true
		dfs(b, i+1)
		b = b[:len(b)-1]
		visit[i] = false
	}
}
