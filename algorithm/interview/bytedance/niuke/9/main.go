package main

/**
合并三个有序数组
*/

func arrayCombine(nums [][]int) []int {
	var n = len(nums)
	var m = (n + 1) / 2
	var mid = n / 2

	for n != 1 {
		for i := 0; i < mid; i++ {
			nums[i] = merge(nums[i], nums[i+m])
		}
		n = m
		mid = n / 2
		m = (n + 1) / 2

	}
	return nums[0]
}

func merge(arr1 []int, arr2 []int) []int {
	return nil
}
