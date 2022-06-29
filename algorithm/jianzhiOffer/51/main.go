package main

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	leftCount := mergeSort(nums, left, mid)
	rightCount := mergeSort(nums, mid+1, right)
	count := leftCount + rightCount
	var i = left
	var j = mid + 1
	var tmp []int
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			count += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		count += right - (mid + 1) + 1
		i++
	}
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}
	for k := left; k <= right; k++ {
		nums[k] = tmp[k-left]
	}
	return count
}
