package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m = len(nums1)
	var n = len(nums2)
	if (m+n)%2 == 0 {
		k1 := (m + n) / 2
		k2 := k1 + 1
		mid1 := float64(findKthElement(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, k1))
		mid2 := float64(findKthElement(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, k2))
		return (mid1 + mid2) / 2
	} else {
		k := (m+n)/2 + 1
		mid := findKthElement(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, k)
		return float64(mid)
	}
}

func findKthElement(nums1 []int, nums2 []int, left1, right1, left2, right2 int, k int) int {
	if left1 > right1 {
		return nums2[left2+k-1]
	}
	if left2 > right2 {
		return nums1[left1+k-1]
	}
	if k == 1 {
		return min(nums1[left1], nums2[left2])
	}
	k1 := min(k/2, right1-left1+1)
	k2 := min(k-k1, right2-left2+1)
	if nums1[left1+k1-1] <= nums2[left2+k2-1] {
		return findKthElement(nums1, nums2, left1+k1, right1, left2, right2, k-k1)
	} else {
		return findKthElement(nums1, nums2, left1, right1, left2+k2, right2, k-k2)
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
