package main

import "fmt"

func main() {
	var nums = []int{2, 2, 1, 1, 0, 2, 9}
	fmt.Println(majorityElement2(nums))
}

/**
构建map
*/
func majorityElement2(nums []int) int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}
	max := 0
	num := 0
	for k, v := range m {
		if v > max && v >= len(nums)/2 {
			num = k
			max = v
		}
	}
	return num
}

/**
分治法
*/
func majorityElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	return getMajorityElement(nums, left, right)
}

func getMajorityElement(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftMajor := getMajorityElement(nums, left, mid)
	rightMajor := getMajorityElement(nums, mid+1, right)
	if leftMajor == rightMajor {
		return leftMajor
	} else {
		leftCount := 0
		rightCount := 0
		for i := left; i < right; i++ {
			if nums[i] == leftMajor {
				leftCount++
			} else if nums[i] == rightMajor {
				rightCount++
			}
		}
		if leftCount >= rightCount {
			return leftMajor
		} else {
			return rightMajor
		}
	}
}
