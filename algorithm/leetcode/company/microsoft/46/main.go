package main

var res [][]int
var visit []bool

func permute(nums []int) [][]int {
	res = nil
	visit = make([]bool, len(nums))
	dfs(nums, []int{})
	return res
}

func dfs(nums []int, b []int) {
	if len(b) == len(nums) {
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visit[i] == false {
			visit[i] = true
			b = append(b, nums[i])
			dfs(nums, b)
			b = b[:len(b)-1]
			visit[i] = false
		}
	}
}
