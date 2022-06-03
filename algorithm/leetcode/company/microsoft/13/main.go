package main

import "fmt"

func main() {
	fmt.Println(romanToInt("LVIII"))
}

func romanToInt2(s string) int {
	var n = len(s)
	var res int
	for i := 0; i < n; i++ {
		if i < n-1 && romToI(string(s[i])) < romToI(string(s[i+1])) {
			res = res - romToI(string(s[i]))
		} else {
			res = res + romToI(string(s[i]))
		}
	}
	return res
}

func romToI(s string) int {
	switch s {
	case "M":
		return 1000
	case "D":
		return 500
	case "C":
		return 100
	case "L":
		return 50
	case "X":
		return 10
	case "V":
		return 5
	case "I":
		return 1
	}
	return 0
}

func romanToInt(s string) int {
	var n = len(s)
	var res int
	var m = make(map[string]int)
	m["M"] = 1000
	m["CM"] = 900
	m["D"] = 500
	m["CD"] = 400
	m["C"] = 100
	m["XC"] = 90
	m["L"] = 50
	m["XL"] = 40
	m["X"] = 10
	m["IX"] = 9
	m["V"] = 5
	m["IV"] = 4
	m["I"] = 1
	for i := 0; i < n; i++ {
		if i < n-1 && m[string(s[i])] < m[string(s[i+1])] {
			res = res - m[string(s[i])]
		} else {
			res = res + m[string(s[i])]
		}
	}
	return res
}
