package main

func minWindow(s string, t string) string {
	var m = len(s)
	var n = len(t)
	if m < n {
		return ""
	}
	var sm = make(map[byte]int)
	var tm = make(map[byte]int)
	for i := 0; i < n; i++ {
		sm[s[i]]++
	}

	for i := 0; i < n; i++ {
		tm[t[i]]++
	}
	var res = ""
	var left = 0 // [left,right)
	var right = n
	var tag bool
	for right <= m {
		tag = true
		for key, _ := range tm {
			if sm[key] < tm[key] {
				tag = false
				break
			}
		}
		if tag {
			if len(res) == 0 || len(res) > right-left {
				res = s[left:right]
			}
			sm[s[left]]--
			left++
		} else if right < m {
			sm[s[right]]++
			right++
		} else {
			right++
		}
	}
	return res
}
