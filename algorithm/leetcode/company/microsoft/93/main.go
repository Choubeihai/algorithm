package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

var ss string
var res []string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	ss = s
	res = nil
	dfs([]string{}, 0)
	return res
}

func dfs(b []string, left int) {
	if left == len(ss) && len(b) == 4 {
		var sb = &strings.Builder{}
		for i := 0; i < 4; i++ {
			sb.WriteString(b[i])
			sb.WriteString(".")
		}
		ans := sb.String()
		res = append(res, ans[:len(ans)-1])
	}
	if left >= len(ss) {
		return
	}
	if len(b) >= 4 {
		return
	}

	for i := 1; i <= len(ss); i++ {
		if i <= left {
			continue
		}
		subString := ss[left:i]
		num, _ := strconv.ParseInt(subString, 10, 32)
		if num > 255 {
			return
		}
		if num != 0 && subString[0] == '0' {
			return
		}
		if num == 0 && len(subString) >= 2 {
			return
		}
		b = append(b, subString)
		dfs(b, i)
		b = b[:len(b)-1]
	}
}
