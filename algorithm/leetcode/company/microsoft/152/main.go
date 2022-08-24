package main

func maxProduct(nums []int) int {
	var n = len(nums)
	var maxDp = make([]int, n)
	var minDp = make([]int, n)
	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	var res = maxDp[0]
	for i := 1; i < n; i++ {
		maxDp[i] = max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i], nums[i])
		minDp[i] = min(minDp[i-1]*nums[i], maxDp[i-1]*nums[i], nums[i])
	}

	for i := 1; i < n; i++ {
		if res < maxDp[i] {
			res = maxDp[i]
		}
	}
	return res

}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
