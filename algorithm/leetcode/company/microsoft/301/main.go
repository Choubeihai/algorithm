package main

// 1. dfs删除左括号和右括号
//2. 判断是否顺序正常

var res []string

func removeInvalidParentheses(s string) []string {
	res = nil
	var left int  // 记录左括号冗余多少
	var right int // 记录有括号冗余多少
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	if left == 0 && right == 0 {
		res = append(res, s)
		return res

	}
	dfs(s, 0, left, right)
	return res

}
func dfs(s string, start int, left, right int) {
	if left == 0 && right == 0 {
		// 判断删除冗余的左括号和右括号后，判断剩下的括号组成是否合理
		var leftCount = 0
		var rightCount = 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				leftCount++
			} else if s[i] == ')' {
				rightCount++
				if rightCount > leftCount {
					return
				}
			}
		}
		if leftCount == rightCount {
			res = append(res, s)
		}
		return
	}
	if start >= len(s) {
		return
	}
	for i := start; i < len(s); i++ {
		if i > start && s[i] == s[i-1] {
			continue
		}
		if s[i] == '(' && left > 0 {
			dfs(s[:i]+s[i+1:], i, left-1, right)
		} else if s[i] == ')' && right > 0 {
			dfs(s[:i]+s[i+1:], i, left, right-1)
		}
	}
}
