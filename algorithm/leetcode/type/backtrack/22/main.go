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
	backtrack(sb, n, n)
	return res
}

func backtrack(b []byte, left, right int) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		combination := strings.ReplaceAll(string(b), "\u0000", "")
		res = append(res, combination)
		return
	}
	if left > 0 {
		left--
		b = append(b, '(')
		backtrack(b, left, right)
		b = b[:len(b)-1]
		left++
	}
	if right > 0 {
		right--
		b = append(b, ')')
		backtrack(b, left, right)
		b = b[:len(b)-1]
		right++
	}
}
