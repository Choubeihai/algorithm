package main

import (
	"fmt"
)

/***
给定一个二进制数组，元素不是0，就是1， 找到一个连续的0和1相等的最大子数组

思路：
将0变成-1，得到一个新的数组。然后构建前缀和sum数组，其中sum[k]对应前k-1个数相加，
遍历sum数组，如果出现sum[i]和前面的sum[j]相等，则说明长度为j-i的连续0和1相等的子数组。
*/

func main() {
	arr := []int{1, 0, 1, 1, 0, 0, 0, 1, 1}
	fmt.Println(findSubInt(arr))
}

func findSubInt(arr []int) []int {
	var sum = make([]int, len(arr)+1)
	var m = make(map[int]int)
	var maxLength = 0
	var left = 0
	sum[0] = 0
	for i := 1; i < len(sum); i++ {
		if arr[i-1] == 0 {
			sum[i] = sum[i-1] + (-1)
		} else {
			sum[i] = sum[i-1] + 1
		}
	}

	m[sum[0]] = -1 //sum[0]对应的前缀为-1
	for i := 1; i < len(sum); i++ {
		if index, ok := m[sum[1]]; ok {
			length := i - 1 - index
			if length > maxLength {
				maxLength = length
				left = index + 1
			}
		} else {
			m[sum[i]] = i - 1
		}
	}
	return arr[left : left+maxLength]
}
