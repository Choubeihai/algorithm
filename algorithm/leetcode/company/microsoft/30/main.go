package main

import "fmt"

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}

// 基本思路：通过map进行比较，没啥难的
func findSubstring(s string, words []string) []int {
	var res []int
	var wordMap = make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordMap[words[i]]++
	}
	var oneWordLen = len(words[0])
	var wordsLen = oneWordLen * len(words)
	for start := 0; start+wordsLen <= len(s); start++ {
		subS := s[start : start+wordsLen]
		m := make(map[string]int)
		i := 0
		for ; i < len(subS); i = i + oneWordLen {
			ss := subS[i : i+oneWordLen]
			m[ss]++
			if m[ss] > wordMap[ss] {
				break
			}
		}
		if i == wordsLen {
			res = append(res, start)
		}
	}
	return res
}
