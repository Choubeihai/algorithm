package main

import "fmt"

/**
字节面试题，求大佬们看看，数组A中给定可以使用的1~9的数，返回由A数组中的元素组成的小于n的最大数。
例如A={1, 2, 4, 9}，x=2533，返回2499
*/

func main() {
	A := []int{1, 2, 4, 9}
	num := 2533
	res := findNum(A, num)
	fmt.Println(res)
}

var res int

func findNum(A []int, num int) int {
	res = 0
	dfs(A, num, 0)
	return res
}

func dfs(A []int, target int, digit int) {
	if digit > target {
		return
	}
	res = max(res, digit)
	for i := 0; i < len(A); i++ {
		digit = digit*10 + A[i]
		dfs(A, target, digit)
		digit = digit / 10
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
