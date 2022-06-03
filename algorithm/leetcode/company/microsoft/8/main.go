package main

import "math"

func myAtoi(s string) int {
	n := len(s)
	index := 0
	for index < n && s[index] == ' ' {
		index++
	}
	if index == n {
		return 0
	}

	sign := 1
	var res int32
	var last int32
	for i := index; i < n; i++ {
		if i == index {
			if s[i] == '-' {
				sign = -1
				continue
			} else if s[i] == '+' {
				continue
			} else if s[i] < '0' || s[i] > '9' {
				return 0
			}
			last = res
			res = res*10 + int32(s[i]-'0')
		} else {
			if s[i] < '0' || s[i] > '9' {
				break
			}
			last = res
			res = res*10 + int32(s[i]-'0')
			if last != res/10 { // 如果不相等，就是溢出了
				if sign == -1 {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
		}
	}
	return sign * int(res)
}
