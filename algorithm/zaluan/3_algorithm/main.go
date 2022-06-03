package main

import "fmt"

/**
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
*/

func main() {
	var s = "我是中国人"
	s2 := reversString(s)
	fmt.Println(s2)
}

func reversString(s string) string {
	str := []rune(s)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-i-1], str[i]
	}
	return string(str)
}
