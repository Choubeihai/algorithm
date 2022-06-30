package main

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	var res = 1
	for n > 0 {
		if n > 4 {
			n -= 3
			res *= 3
		} else if n == 4 {
			n -= 4
			res *= 4
		} else if n == 3 {
			n -= 3
			res *= 3
		} else {
			res *= n
			n = 0
		}
	}
	return res
}
