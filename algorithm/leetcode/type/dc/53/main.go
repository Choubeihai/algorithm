package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray2(nums))
}

func maxSubArray(nums []int) int {
	var left = 0
	var right = len(nums) - 1
	return getMaxContinuousSum(nums, left, right)
}

/**
  获取连续子串的最大值
  使用分治算法
  易分难合
*/
func getMaxContinuousSum(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftMaxContinuousSum := getMaxContinuousSum(nums, left, mid)     // 左子串连续最大值
	rightMaxContinuousSum := getMaxContinuousSum(nums, mid+1, right) // 右子串连续最大值

	sum := 0
	leftSum := math.MinInt
	rightSum := math.MinInt
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if leftSum < sum {
			leftSum = sum
		}
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if rightSum < sum {
			rightSum = sum
		}
	}
	var nowMaxContinuousSum int
	if leftSum > 0 {
		nowMaxContinuousSum += leftSum
	}
	if rightSum > 0 {
		nowMaxContinuousSum += rightSum
	}
	return max(leftMaxContinuousSum, nowMaxContinuousSum, rightMaxContinuousSum)
}

/**
使用动态规划
*/
func maxSubArray2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	var sum = make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = max1(sum[i-1]+nums[i], nums[i])
	}
	var maxSum = sum[0]
	for i := 1; i < n; i++ {
		if maxSum < sum[i] {
			maxSum = sum[i]
		}
	}
	return maxSum
}

func max1(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}
