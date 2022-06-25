package main

func isScramble(s1 string, s2 string) bool {
	m := len(s1)
	n := len(s2)
	if m != n {
		return false
	}
	// dp[i][j][n+1]:
	var dp = make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
	}

	// 枚举总长度length
	for length := 2; length <= n; length++ {
		// s1的起点位置index
		for i := 0; i <= n-length; i++ {
			// s2的起点位置index
			for j := 0; j <= n-length; j++ {
				// 枚举分割点index
				for k := 1; k <= length-1; k++ {
					// s11——> s21, s12——> s22
					if dp[i][j][k] && dp[i+k][j+k][length-k] {
						dp[i][j][length] = true
						break
					}
					//  s11——> s22, s11——> s22
					if dp[i][j+length-k][k] && dp[i+k][j][length-k] {
						dp[i][j][length] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
