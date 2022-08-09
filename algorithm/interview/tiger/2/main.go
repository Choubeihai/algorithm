package main

import (
	"fmt"
)

/**
有一个二进制字符串，字符串循环右移k位，然后得到一个新的二进制字符串，求这个字符串的十进制是多少。
*/

func main() {
	s := "10110"
	k := -2
	fmt.Println(rightMove(s, k))
}

func rightMove(s string, k int) int {
	n := len(s)
	var res int
	if k == 0 {
		k = k % n
		bit := 1
		for i := n - 1; i >= 0; i-- {
			num := int(s[i] - '0')
			res = res + num*bit
			bit = bit * 2
		}
		return res
	} else if k < 0 {
		k = -k
		k = k % n
		bit := 1
		for i := k - 1; i >= 0; i-- {
			num := int(s[i] - '0')
			res = res + num*bit
			bit = bit * 2
		}
		for i := n - 1; i >= k; i-- {
			num := int(s[i] - '0')
			res = res + num*bit
			bit = bit * 2
		}
		return res
	} else { // k > 0
		k = k % n
		bit := 1
		for i := n - k - 1; i >= 0; i-- {
			num := int(s[i] - '0')
			res = res + num*bit
			bit = bit * 2
		}
		fmt.Println()
		for i := n - 1; i >= n-k; i-- {
			num := int(s[i] - '0')
			res = res + num*bit
			bit = bit * 2
		}
		return res
	}
}
