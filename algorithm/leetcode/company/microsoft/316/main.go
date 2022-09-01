package main

func removeDuplicateLetters(s string) string {
	var n = len(s)
	var stack = make([]int, n)
	var p = -1

	var m = make(map[byte]int)
	var visit = make(map[byte]bool)

	for i := 0; i < n; i++ {
		m[s[i]]++
	}

	for i := 0; i < n; i++ {
		if visit[s[i]] {
			m[s[i]]--
			continue
		}

		for p >= 0 && m[s[stack[p]]] > 0 && s[i] < s[stack[p]] {
			cur := stack[p]
			p--
			visit[s[cur]] = false
		}
		p++
		stack[p] = i
		visit[s[i]] = true
		m[s[stack[p]]]--
	}

	var res = ""
	for i := 0; i <= p; i++ {
		res += string(s[stack[i]])
	}
	return res
}
