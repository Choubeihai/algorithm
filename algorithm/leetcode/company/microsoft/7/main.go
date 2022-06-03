package main

import (
	"fmt"
	"math"
)

func main() {
	x := -123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	flag := 1
	if x < 0 {
		flag = -1
		if x <= math.MinInt32 {
			return 0
		}
		x = -x
	} else {
		if x > math.MaxInt32+1 {
			return 0
		}
	}
	res := 0
	for x != 0 {
		remainder := x % 10
		res = res*10 + remainder
		x = x / 10
	}
	res = flag * res
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	} else {
		return res
	}
}
