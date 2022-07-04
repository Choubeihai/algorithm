package main

import "fmt"

func main() {
	var x float64 = 2
	var n = 10
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var flag = 1
	if n < 0 {
		flag = -1
		n = -n
	}
	var res float64 = 1
	var m = 1
	var tmpRes = x
	for n != 0 {
		if n > m {
			n = n - m
		} else {
			m = 1
			tmpRes = x
			n = n - m
		}
		res = res * tmpRes
		tmpRes = tmpRes * tmpRes
		m = m * 2
	}
	if flag == 1 {
		return res
	} else {
		return 1 / res
	}
}
