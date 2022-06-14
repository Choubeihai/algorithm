package main

import "fmt"

func main() {
	digits := "2"
	fmt.Println(letterCombinations(digits))

}

// 回溯
var myDigits string
var s []string
var res []string

func letterCombinations(digits string) []string {
	// leetcode不会对外部变量进行初始化，这是leetcode的原因
	myDigits = digits
	s = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res = nil
	if len(myDigits) == 0 {
		return nil
	}
	combination := make([]byte, 0, len(digits))
	backtrack(combination, 0)
	return res
}

func backtrack(combination []byte, index int) {
	if index == len(myDigits) {
		res = append(res, string(combination))
		return
	}
	if index > len(myDigits) {
		return
	}
	ss := s[myDigits[index]-'0']
	for i := 0; i < len(ss); i++ {
		combination = append(combination, ss[i])
		backtrack(combination, index+1)
		combination = combination[:len(combination)-1]
	}
}
