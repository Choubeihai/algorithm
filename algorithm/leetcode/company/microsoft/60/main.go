package main

import "strconv"

/**
一开始的思路是使用全排列做，但是时间超时了。
所以换种思路
*/

func getPermutation(n int, k int) string {
	var fact = make([]int, n+1) // 记录阶乘
	fact[0] = 1                 // 0的阶乘是1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i
	}
	var res string
	var used = make([]bool, n+1) // 表明哪些数字被使用过

	for i := 1; i <= n; i++ { // 欲选的排列中还有几个数字没被确定
		cnt := fact[n-i]
		for j := 1; j <= n; j++ { // 判断数字
			if !used[j] {
				if k > cnt {
					k = k - cnt
				} else {
					res += strconv.Itoa(j)
					used[j] = true
					break
				}
			}
		}
	}
	return res
}
