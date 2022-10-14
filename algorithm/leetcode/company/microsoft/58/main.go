package main

import "fmt"

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	res := 0
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if i > 0 && s[i-1] == ' ' {
				continue
			} else {
				res = count
				count = 0
			}
		} else {
			count++
		}
	}

	if count != 0 {
		res = count
	}
	return res

}
