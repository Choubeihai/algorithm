package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 3
	res = generateParenthesis(n)
	fmt.Println(res)
}

var res []string

func generateParenthesis(n int) []string {
	res = nil
	sb := make([]byte, n*2)
	recur(sb, n, n)
	return res
}

func recur(sb []byte, left, right int) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		combination := strings.ReplaceAll(string(sb), "\u0000", "")
		res = append(res, combination)
		return
	}
	if left > 0 {
		left--
		sb = append(sb, '(')
		recur(sb, left, right)
		sb = sb[:len(sb)-1]
		left++
	}
	if right > 0 {
		right--
		sb = append(sb, ')')
		recur(sb, left, right)
		sb = sb[:len(sb)-1]
		right++
	}
}
