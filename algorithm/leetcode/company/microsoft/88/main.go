package main

// 不申请额外的空间
func merge1(nums1 []int, m int, nums2 []int, n int) {
	var i = m - 1
	var j = n - 1
	var k = m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	num11 := make([]int, m)
	copy(num11, nums1[:m])
	var i int
	var j int
	var k int
	for i < m && j < n {
		if num11[i] < nums2[j] {
			nums1[k] = num11[i]
			k++
			i++
		} else {
			nums1[k] = nums2[j]
			k++
			j++
		}
	}
	for i < m {
		nums1[k] = num11[i]
		k++
		i++
	}
	for j < n {
		nums1[k] = nums2[j]
		k++
		j++
	}
}
