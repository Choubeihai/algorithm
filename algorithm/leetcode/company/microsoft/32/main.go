package main

import "fmt"

func main() {
	s := ")()())"
	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	var res int
	n := len(s)
	stack := make([]int, n+1)
	p := -1

	stack[0] = -1
	p++

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			p++
			stack[p] = i
		} else { // s[i] == ')'
			p--
			if p == -1 {
				p++
				stack[p] = i
			} else {
				res = max(res, i-stack[p])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
