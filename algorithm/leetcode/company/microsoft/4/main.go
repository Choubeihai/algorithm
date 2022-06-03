package main

import (
	"fmt"
)

func main() {
	nums1 := []int{2, 3, 4}
	nums2 := []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		midIndex1 := n/2 - 1
		midIndex2 := n / 2
		kth1 := findKth2(nums1, nums2, midIndex1+1)
		kth2 := findKth2(nums1, nums2, midIndex2+1)
		return float64(kth1+kth2) / 2
	} else {
		midIndex := n / 2
		kth := findKth2(nums1, nums2, midIndex+1)
		return float64(kth)
	}
}

// 内部循环 寻找第k大元素
func findKth2(num1 []int, num2 []int, k int) int {
	index1 := 0
	index2 := 0
	for {
		if len(num1) == index1 {
			return num2[index2+k-1]
		}
		if len(num2) == index2 {
			return num1[index1+k-1]
		}
		if k == 1 {
			return min(num1[index1], num2[index2])
		}

		k1 := k / 2
		k2 := k - k1
		newIndex1 := min(index1+k1-1, len(num1)-1)
		newIndex2 := min(index2+k2-1, len(num2)-1)
		if num1[newIndex1] <= num2[newIndex2] {
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

// 找到第K大元素 递归
func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	k1 := min(k/2, len(nums1))  // 保证不超过num1的长度
	k2 := min(k-k1, len(nums2)) // 保证不超过num2的长度
	if nums1[k1-1] <= nums2[k2-1] {
		return findKth(nums1[k1:], nums2, k-k1)
	} else {
		return findKth(nums1, nums2[k2:], k-k2)
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
