package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	var m = make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		bs := make([]int, 26) // 记录字母出现的次数
		for j := 0; j < len(strs[i]); j++ {
			index := strs[i][j] - 'a'
			bs[index]++
		}
		byteBuffer := &bytes.Buffer{}
		for j := 0; j < len(bs); j++ {
			if bs[j] != 0 {
				byteBuffer.WriteByte(byte('a' + j))
				byteBuffer.WriteString(strconv.Itoa(bs[j]))
			}
		}
		key := byteBuffer.String()
		if value, ok := m[key]; ok {
			value = append(value, strs[i])
			m[key] = value
		} else {
			value = []string{strs[i]}
			m[key] = value
		}
	}
	var res [][]string
	for _, value := range m {
		res = append(res, value)
	}
	return res
}
