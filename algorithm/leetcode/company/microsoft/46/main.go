package main

import "fmt"

var flag []bool
var myNum []int
var res [][]int

func main() {
	nums := []int{1, 2, 3, 4}
	//fmt.Println(permute1(nums))
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	flag = make([]bool, 21) // 因为nums中的数字在[-10, 10]
	myNum = nums
	res = nil
	dfs([]int{})
	return res
}

func dfs(b []int) {
	if len(b) == len(myNum) {
		c := make([]int, len(b))
		copy(c, b) // 一定要拷贝一份，否则会出错
		res = append(res, c)
		return
	}

	for i := 0; i < len(myNum); i++ {
		if !flag[myNum[i]+10] {
			flag[myNum[i]+10] = true
			b = append(b, myNum[i])
			dfs(b)
			flag[myNum[i]+10] = false
			b = b[:len(b)-1]
		}
	}
}
