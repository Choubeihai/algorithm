package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	l, r := 0, x
	var res int
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return res
}
