package main

import "fmt"

func main() {
	s := "{}[]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	stack := make([]byte, n)
	p := -1
	for i := 0; i < n; i++ {
		switch s[i] {
		case ')':
			if p == -1 {
				return false
			} else {
				pair := stack[p]
				if pair != '(' {
					return false
				}
				p--
			}
		case ']':
			if p == -1 {
				return false
			} else {
				pair := stack[p]
				if pair != '[' {
					return false
				}
				p--
			}
		case '}':
			if p == -1 {
				return false
			} else {
				pair := stack[p]
				if pair != '{' {
					return false
				}
				p--
			}
		default:
			p++
			stack[p] = s[i]
		}
	}
	if p != -1 {
		return false
	} else {
		return true
	}
}
