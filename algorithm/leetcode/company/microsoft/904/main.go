package main

/**
双指针问题
*/

func totalFruit(fruits []int) int {
	m := make(map[int]int)
	left := 0
	right := 0
	res := 0
	for right < len(fruits) {
		ttype := fruits[right]
		m[ttype]++
		if len(m) > 2 {
			ttype = fruits[left]
			m[ttype]--
			if m[ttype] == 0 {
				delete(m, ttype)
			}
			left++
		} else {
			if right-left+1 > res {
				res = right - left + 1
			}
		}
		right++

	}
	return res
}
