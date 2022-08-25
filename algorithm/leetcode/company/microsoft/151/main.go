package main

import "strings"

func reverseWords(s string) string {
	var start = 0
	var end = 0
	var res = &strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		} else {
			end = i + 1
			for i >= 0 && s[i] != ' ' {
				i--
			}
			start = i + 1
			res.WriteString(s[start:end])
			res.WriteByte(' ')
		}
	}
	ans := res.String()
	ans = ans[:len(ans)-1]
	return ans
}
