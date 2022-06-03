package main

import "fmt"

func main() {
	text := "hello"
	pattern := "abaabcac"
	fmt.Println(strStr(text, pattern))
}

// KMP算法
// 寻找最长公共前缀和最长公共后缀
func strStr(text, pattern string) int {
	if text == "" {
		return -1
	}
	if pattern == "" {
		return 0
	}

	// 构建共有字符串表
	var next = make([]int, len(pattern))
	next[0] = 0 // 规定
	var i = 1
	var length = 0 // 上次共有字符串的长度(即最大前缀和最大后缀相同时候的长度)
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			next[i] = length
			i++
		} else {
			if length > 0 {
				length = next[length-1]
			} else { // length=0
				next[i] = length
				i++
			}
		}
	}
	// 移动一位
	for j := len(next) - 1; j > 0; j-- {
		next[j] = next[j-1]
	}
	next[0] = -1

	i = 0
	j := 0
	for i < len(text) && j < len(pattern) {
		if text[i] != pattern[j] {
			j = next[j]
			if j == -1 {
				j++
				i++
			}
		} else {
			j++
			i++
		}
	}
	if j == len(pattern) {
		return i - j
	} else {
		return -1
	}

}
