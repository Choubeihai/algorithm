package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	var res *bytes.Buffer
	for i := 1; i <= n; i++ {
		if i == 1 {
			res = &bytes.Buffer{}
			res.WriteString("1")
		} else {
			tmpStr := res.String()
			res = &bytes.Buffer{}
			count := 1
			c := tmpStr[0] - '0'
			for j := 1; j < len(tmpStr); j++ {
				if tmpStr[j] == tmpStr[j-1] {
					count++
				} else {
					res.WriteString(strconv.FormatInt(int64(count), 10))
					res.WriteString(strconv.FormatInt(int64(c), 10))
					count = 1
					c = tmpStr[j] - '0'
				}
			}
			res.WriteString(strconv.FormatInt(int64(count), 10))
			res.WriteString(strconv.FormatInt(int64(c), 10))
		}
	}
	return res.String()
}
