package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	start := 0
	res := 0
	for i := 0; i < len(s); i++ {
		oldIdx, ok := m[s[i]]
		if ok && start < oldIdx+1 {
			start = oldIdx + 1
		}
		m[s[i]] = i
		if i-start+1 > res {
			res = i - start + 1
		}
	}
	return res
}
