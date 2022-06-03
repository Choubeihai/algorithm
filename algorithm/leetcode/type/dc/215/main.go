package main

import "fmt"

/**
第k大的元素，在倒序排序下，index是k-1
*/

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums)
	return find(nums, left, right-1, k)
}

func find(nums []int, left, right, k int) int {
	part := partition(nums, left, right)
	if part == k-1 {
		return nums[part]
	} else if part < k-1 {
		return find(nums, part+1, right, k)
	} else {
		return find(nums, left, part-1, k)
	}
}

func partition(nums []int, left, right int) int {
	i := left
	j := right
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] <= pivot {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i] > pivot {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = pivot
	return i
}
