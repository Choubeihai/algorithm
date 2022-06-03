package main

import "fmt"

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	chars := []byte(s)
	length := len(chars)
	if length == 0 || length == 1 {
		return length
	}
	m := make(map[byte]int)
	start := 0
	res := 0
	for i := 0; i < length; i++ {
		if latestIndex, ok := m[chars[i]]; ok {
			m[chars[i]] = i
			if start < latestIndex+1 {
				start = latestIndex + 1
			}
			res = max(res, i-start+1)
		} else {
			m[chars[i]] = i
			res = max(res, i-start+1)
		}
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
