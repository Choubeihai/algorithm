package main

import "fmt"

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	var count = 0
	var res int
	var n = len(s)
	var index = 0
	for index < n {
		if s[index] == ' ' {
			if index > 0 && s[index-1] == ' ' {
				index++
			} else {
				res = count
				count = 0
				index++
			}
		} else {
			count++
			index++
		}
	}
	if count != 0 {
		res = count
	}
	return res
}
