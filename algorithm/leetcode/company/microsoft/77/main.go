package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

var nn int
var res [][]int
var kk int
var visit []bool

func combine(n int, k int) [][]int {
	nn = n
	res = nil
	kk = k
	visit = make([]bool, nn+1)
	dfs([]int{}, 0)
	return res
}

func dfs(b []int, index int) {
	if len(b) == kk {
		c := make([]int, kk)
		copy(c, b)
		res = append(res, c)
		return
	} else if len(b) > kk {
		return
	}
	for i := 1; i <= nn; i++ {
		if !visit[i] && index <= i {
			visit[i] = true
			b = append(b, i)
			dfs(b, i+1)
			visit[i] = false
			b = b[:len(b)-1]
		}
	}
}
