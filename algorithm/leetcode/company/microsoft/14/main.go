package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	str := strs[0]
	res := &strings.Builder{}
	for i := 0; i < len(str); i++ {
		for j := 1; j < len(strs); j++ {
			if i < len(strs[j]) && strs[j][i] == str[i] {
				continue
			} else {
				return res.String()
			}
		}
		res.WriteByte(str[i])
	}
	return res.String()
}
