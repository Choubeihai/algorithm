package main

import "fmt"

func main() {
	s := "b"
	p := "a*"
	fmt.Println(isMatch(s, p))
}

// 和第 44 题很像 115
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) //dp[m+1][n+1] 表示s的前i个字符与p的前j个字符是否能够匹配
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	//当s的长度和p的长度都为0时，则二者可以匹配
	dp[0][0] = true

	// 当p的长度为0时，s的长度不为0时，则均设置为false
	for i := 1; i <= m; i++ {
		dp[i][0] = false
	}

	// 	// 当s的长度为0时，p的长度不为0时，因为*可以匹配0个字符，所以不能一刀切设置为false，所以有可能为true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2] // // 因为第i个字符为*，dp[0][i]取决于dp[0][i-2]
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
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					// dp[i][j-2]表示*代表0个字符，所以取dp[i][j-2]
					// dp[i-1][j]表示*代表1个或者1个以上字符，所以将s的字符删除一个，仍然可以匹配，取dp[i-1][j]
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else { // // 模式串*的前一个字符不能够跟文本串的末位匹配
					dp[i][j] = dp[i][j-2] // *只能匹配0次
				}

			}
		}
	}
	return dp[m][n]
}
