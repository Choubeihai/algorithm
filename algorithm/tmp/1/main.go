package main

import "fmt"

func main() {
	nums := []int{1, 2, 5}
	sum := 100
	numsOfCombination(nums, sum)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
	fmt.Println(cnt)
}

var res [][]int
var cnt int

func numsOfCombination(nums []int, sum int) ([][]int, int) {
	res = nil
	dfs(nums, 0, []int{}, sum)
	return res, cnt
}

func dfs(nums []int, index int, b []int, sum int) {
	if sum == 0 {
		cnt++
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
		return
	}

	if index >= len(nums) {
		return
	}

	if sum < 0 {
		return
	}

	for i := index; i < len(nums); i++ {
		b = append(b, nums[i])
		sum -= nums[i]
		dfs(nums, i, b, sum)
		sum += nums[i]
		b = b[:len(b)-1]
	}
}
