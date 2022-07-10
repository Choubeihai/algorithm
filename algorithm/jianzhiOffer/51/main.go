package main

func reversePairs(nums []int) int {
	return sort(nums, 0, len(nums)-1)
}

func sort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	leftCount := sort(nums, left, mid)
	rightCount := sort(nums, mid+1, right)
	return leftCount + rightCount + merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) int {
	arr1 := make([]int, mid-left+1)
	arr2 := make([]int, right-mid)
	copy(arr1, nums[left:mid+1])
	copy(arr2, nums[mid+1:])
	var i int
	var j int
	var k = left
	var count int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			nums[k] = arr1[i]
			i++
			k++
		} else {
			nums[k] = arr2[j]
			count += len(arr1) - 1 - i + 1
			j++
			k++
		}
	}
	for i < len(arr1) {
		nums[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		nums[k] = arr2[j]
		j++
		k++
	}
	return count
}
