package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var n = len(nums2)
	var m = make(map[int]int)
	var stack = make([]int, n)
	var p = -1
	for i := 0; i < n; i++ {
		for p >= 0 && nums2[stack[p]] < nums2[i] {
			cur := stack[p]
			p--
			m[nums2[cur]] = nums2[i]
		}
		p++
		stack[p] = i
	}
	var res = make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if v, ok := m[nums1[i]]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}
