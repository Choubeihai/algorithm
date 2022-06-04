package string

import "fmt"

// 遍历字符串
func traverse(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // uint8,即byte
	}
	for _, r := range s {
		fmt.Println(r) // int32,即rune
	}
}
