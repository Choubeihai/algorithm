package main

// 方法一 笨法子
func nextGreaterElements(nums []int) []int {
	var n = len(nums)
	var res = make([]int, n)
	var tag bool
	for i := 0; i < n; i++ {
		tag = false
		for j := (i + 1) % n; j != i; j = (j + 1) % n {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				tag = true
				break
			}
		}
		if !tag {
			res[i] = -1
		}
	}
	return res
}

// 方法二： 单调栈
func nextGreaterElements2(nums []int) []int {
	var n = len(nums)
	var res = make([]int, n)
	var stack = make([]int, 2*n-1)
	var p = -1
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	for i := 0; i < 2*n-1; i++ {
		index := i % n
		for p >= 0 && nums[stack[p]] < nums[index] {
			cur := stack[p]
			res[cur] = nums[index]
			p--
		}
		p++
		stack[p] = index
	}
	return res
}
