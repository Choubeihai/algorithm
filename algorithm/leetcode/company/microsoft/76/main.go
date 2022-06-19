package main

import "math"

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	// 窗口[left, right) 左闭右开区间
	var left = 0
	var right = 0
	var start = 0
	var minWindowSize = math.MaxInt32

	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			m[s[right]]--
		}
		right++

		for check(m) {
			if right-left < minWindowSize {
				minWindowSize = right - left
				start = left
			}

			// 左指针 向右移动
			if _, ok := m[s[left]]; ok {
				m[s[left]]++
			}
			left++
		}
	}
	if minWindowSize == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+minWindowSize]
	}
}

func check(m map[byte]int) bool {
	for _, value := range m {
		// 注意这里的value是可以为负数的，为负数的情况就是，相同的字符右
		// 指针扫描的要比t中的多，比如t是"ABC"，窗口中的字符是"ABBC"
		if value > 0 {
			return false
		}
	}
	return true
}
