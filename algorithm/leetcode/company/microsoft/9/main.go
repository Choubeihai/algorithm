package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	y := x
	res := 0
	for y != 0 {
		remainder := y % 10
		res = res*10 + remainder
		y = y / 10
	}
	if res == x {
		return true
	} else {
		return false
	}
}
