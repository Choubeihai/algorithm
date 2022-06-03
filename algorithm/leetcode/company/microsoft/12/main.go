package main

import (
	"fmt"
	"strings"
)

func main() {
	num := 3
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}                    // 罗马数字总共13个数字
	var rom = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"} // 对应的罗马数字
	var res = &strings.Builder{}
	for i := 0; i < len(rom); i++ {
		for num >= nums[i] {
			res.WriteString(rom[i])
			num = num - nums[i]
		}

		if num == 0 {
			break
		}
	}
	return res.String()
}
