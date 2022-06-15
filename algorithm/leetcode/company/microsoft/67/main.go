package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1"
	b := "111"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	var i = len(a) - 1
	var j = len(b) - 1
	var bit int
	var res string
	var index = len(res) - 1
	for i >= 0 && j >= 0 {
		sum := bit + int(a[i]-'0') + int(b[j]-'0')
		bit = sum / 2
		sum = sum % 2
		res = strconv.Itoa(sum) + res
		index--
		i--
		j--
	}
	for i >= 0 {
		sum := bit + int(a[i]-'0')
		bit = sum / 2
		sum = sum % 2
		res = strconv.Itoa(sum) + res
		index--
		i--
	}
	for j >= 0 {
		sum := bit + int(b[j]-'0')
		bit = sum / 2
		sum = sum % 2
		res = strconv.Itoa(sum) + res
		index--
		j--
	}
	if bit != 0 {
		res = strconv.Itoa(bit) + res
		return res
	} else {
		return res
	}
}
