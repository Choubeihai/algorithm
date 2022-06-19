package main

// 很有意义的一道题
func sortColors(nums []int) {
	partition(nums)
}

/**
借助快速排序partition思想 设置循环不变量
[p1 p] 全是0 闭区间
（p p2）全是1 开区间
[p2 len-1] 全是2 闭区间
代码开始之初保证上述三个区间的长度全是0
*/
func partition(nums []int) {
	var p1 = -1
	var p = 0
	var p2 = len(nums)

	for p < p2 {
		if nums[p] == 0 {
			p1++
			nums[p1], nums[p] = nums[p], nums[p1]
			p++
		} else if nums[p] == 1 {
			p++
		} else {
			// nums[p] == 2
			p2--
			nums[p2], nums[p] = nums[p], nums[p2]
		}
	}
}
