package main

import (
	"fmt"
	"strconv"
)

/**
1. 根据逆波兰表达式求值， leetcode 150
2. 根据中缀表达式转化为逆波兰表达式（后缀表达式）， leetcode中无这道题，但牛客网中有人反应这是字节的某道面试题
*/

func main() {
	s := []string{"(", "2", "+", "1", ")", "*", "3"}
	fmt.Println(getReversePolishExpression(s))
}

// 1. 根据逆波兰表达式求值
func evalRPN(tokens []string) int {
	var n = len(tokens)
	var stack = make([]int, n)
	var p = -1

	for i := 0; i < n; i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			v1 := stack[p]
			p--
			v2 := stack[p]
			p--
			if tokens[i] == "+" {
				p++
				stack[p] = v2 + v1
			} else if tokens[i] == "-" {
				p++
				stack[p] = v2 - v1
			} else if tokens[i] == "*" {
				p++
				stack[p] = v2 * v1
			} else {
				p++
				stack[p] = v2 / v1
			}
		} else {
			v, _ := strconv.Atoi(tokens[i])
			p++
			stack[p] = v
		}
	}
	return stack[p]
}

// 2. 根据中序表达式求后序表达式
func getReversePolishExpression(s []string) []string {
	var n = len(s)
	var stack = make([]byte, n)
	var p = -1
	var res []string

	for i := 0; i < n; i++ {
		if s[i] == "(" {
			p++
			stack[p] = '('
		} else if s[i] == ")" {
			for p >= 0 {
				if stack[p] == '(' {
					p--
					break
				} else {
					res = append(res, string(stack[p]))
					p--
				}
			}
		} else if s[i] == "+" || s[i] == "-" {
			for p >= 0 {
				if stack[p] == '/' || stack[p] == '*' {
					res = append(res, string(stack[p]))
					p--
				} else {
					break
				}
			}
			p++
			if s[i] == "+" {
				stack[p] = '+'
			} else {
				stack[p] = '-'
			}
		} else if s[i] == "/" || s[i] == "*" {
			p++
			if s[i] == "*" {
				stack[p] = '*'
			} else {
				stack[p] = '/'
			}
		} else { // 数字
			res = append(res, s[i])
		}
	}
	for p >= 0 {
		res = append(res, string(stack[p]))
		p--
	}
	return res
}
