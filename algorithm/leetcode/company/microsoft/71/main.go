package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	var stack = make([]string, len(path))
	var p = -1
	ss := strings.Split(path, "/")
	for i := 0; i < len(ss); i++ {
		if ss[i] == ".." {
			if p == -1 {
				continue
			} else {
				p--
			}
		} else if ss[i] == "" || ss[i] == "." {
			continue
		} else { // 真实的路径
			p++
			stack[p] = ss[i]
		}
	}
	res := "/" + strings.Join(stack[:p+1], "/")
	return res
}
