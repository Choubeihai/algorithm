package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n) // dp[i][j]: 左边界为i，右边界为j
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	maxLength := 1 // 一定要等于1
	start := 0
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n; i++ { // 右边界
		for j := 0; j < i; j++ { // 左边界
			if s[j] != s[i] {
				dp[j][i] = false
			} else {
				if i-j <= 2 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}

			if dp[j][i] && (i-j+1) > maxLength {
				start = j
				maxLength = i - j + 1
			}
		}
	}
	return s[start : start+maxLength]
}
