package main

import (
	"fmt"
	"math"
)

func main() {
	dividend := 7
	divisor := -3
	fmt.Println(divide(dividend, divisor))
}

// 分治法
func divide(dividend int, divisor int) int {
	sign := 1
	res := 0
	if (dividend^divisor)>>31&1 == 1 { // 按位异或 相同为0 不同为1
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	for dividend >= divisor {
		i := 1
		tmp := divisor
		for dividend >= tmp {
			dividend = dividend - tmp
			res += i
			tmp = tmp << 1
			i = i << 1
		}
	}
	res = res * sign
	if res > math.MaxInt32 || res < math.MinInt32 {
		return math.MaxInt32
	}
	return res
}
