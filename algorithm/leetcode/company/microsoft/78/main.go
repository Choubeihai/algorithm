package main

var myNums []int
var res [][]int

func subsets(nums []int) [][]int {
	myNums = nums
	res = nil
	dfs([]int{}, 0)
	return res
}

func dfs(b []int, index int) {

	if index >= len(myNums) {
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
		return
	}

	c := make([]int, len(b))
	copy(c, b)
	res = append(res, c)

	for i := 0; i < len(myNums); i++ {
		if index > i {
			continue
		}
		b = append(b, myNums[i])
		dfs(b, i+1)
		b = b[:len(b)-1]
	}
}
