package main

/**
判断一个字符串里面的字符都是不相同的。
*/

import "fmt"

func main() {
	s := "ABCHI"
	f := isUniqueString2(s)
	fmt.Println(f)
}

func isUniqueString2(s string) bool {
	var pmask *uint64
	var mask1, mask2, mask3, mask4 uint64
	for _, c := range s {
		n := uint64(c)
		if n < 64 {
			pmask = &mask1 // 指向mask1
		} else if n < 128 {
			pmask = &mask2
			n = n - 64
		} else if n < 192 {
			pmask = &mask3
			n = n - 128
		} else if n < 256 {
			pmask = &mask4
			n = n - 192
		}
		if *pmask&(1<<n) != 0 {
			return false
		}
		*pmask = *pmask | (1 << n)
	}
	return true
}
