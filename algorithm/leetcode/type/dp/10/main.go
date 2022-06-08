package main

import "fmt"

func main() {
	s := "b"
	p := "a*"
	fmt.Println(isMatch(s, p))
}

// 和第 44 题很像
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) //dp[m+1][n+1] 表示s的前i个字符与p的前j个字符是否能够匹配
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true //当s的长度和p的长度都为0时，则二者可以匹配

	for i := 1; i <= m; i++ {
		dp[i][0] = false // 当p的长度为0时，s的长度不为0时，则设置为false
	}

	// s为空，p不为空，由于*可以匹配0个字符，所以有可能为true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		} else {
			dp[0][i] = false
		}
	}

	// 填格子
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 模式串*的前一个字符能够跟文本串的末位匹配上
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j] // 其中，dp[i][j-2]表示 *匹配0次的情况；dp[i-1][j]表示*匹配1次或多次的情况
				} else { // 模式串*的前一个字符不能够跟文本串的末位匹配
					dp[i][j] = dp[i][j-2] // *只能匹配0次
				}
			}
		}
	}
	return dp[m][n]
}
