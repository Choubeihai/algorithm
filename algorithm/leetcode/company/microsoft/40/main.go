package main

import "sort"

func main() {

}

var myCandidates []int
var flag []bool
var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	myCandidates = candidates
	flag = make([]bool, len(myCandidates))
	res = nil
	dfs([]int{}, 0, target)
	return res
}

func dfs(b []int, index int, target int) {
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
		if i > 0 && myCandidates[i] == myCandidates[i-1] && !flag[i-1] {
			continue
		}

		if !flag[i] {
			flag[i] = true
			target = target - myCandidates[i]
			b = append(b, myCandidates[i])
			dfs(b, i, target)
			b = b[:len(b)-1]
			target = target + myCandidates[i]
			flag[i] = false
		}
	}
}
