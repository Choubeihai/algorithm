package main

var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = nil
	dfs(k, n, 1, []int{})
	return res
}

func dfs(k int, sum int, value int, b []int) {
	if k == 0 && sum == 0 {
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
	}
	if k <= 0 {
		return
	}
	if sum <= 0 {
		return
	}
	for i := value; i <= 9; i++ {
		sum -= i
		k--
		b = append(b, i)
		dfs(k, sum, i+1, b)
		sum += i
		k++
		b = b[:len(b)-1]
	}
}
