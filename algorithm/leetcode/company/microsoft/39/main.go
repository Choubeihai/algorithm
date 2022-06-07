package main

import "sort"

var myCandidates []int
var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	myCandidates = candidates
	sort.Ints(myCandidates)
	res = nil
	dfs([]int{}, target, 0)
	return res
}

func dfs(b []int, target int, index int) {
	if target < 0 {
		return
	}
	if target == 0 {
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
		return
	}
	for i := 0; i < len(myCandidates); i++ {
		if index > i {
			continue
		}
		target = target - myCandidates[i]
		b = append(b, myCandidates[i])
		dfs(b, target, i)
		target = target + myCandidates[i]
		b = b[:len(b)-1]
	}
}
