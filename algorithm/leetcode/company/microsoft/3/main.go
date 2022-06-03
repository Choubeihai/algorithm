package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := 0
	length := 0
	for i, ch := range s {
		if index, ok := m[ch]; ok {
			if start < index+1 {
				start = index + 1
			}
			if length < i-start+1 {
				length = i - start + 1
			}
			m[ch] = i
		} else {
			m[ch] = i
			if length < i-start+1 {
				length = i - start + 1
			}
		}
	}
	return length
}
