package main

import (
	"fmt"
)

/**
4,4,6,7,0,1,2
思路：我们将数组一分为二，肯定有一边是有序的，而另一边是可能有序的。
那么我只需要将target与有序的一边比较，如果target位于这一边，则去对这一边进一步一分为二查找；
如果target不位于这一边，则去对另一边(可能不为有序的一边)进行一分为二的查找。
所以说，我们每次都是与有序的一边的进行比较。
*/

func main() {
	var nums = []int{1}
	target := 0
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[left] == target {
		return left
	}
	if nums[mid] == target {
		return mid
	}
	if nums[right] == target {
		return right
	}

	if nums[left] < nums[mid] { // 左边有序
		if target > nums[left] && target < nums[mid] {
			return binarySearch(nums, left+1, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right-1, target)
		}
	} else {
		if target > nums[mid] && target < nums[right] {
			return binarySearch(nums, mid+1, right-1, target)
		} else {
			return binarySearch(nums, left+1, mid-1, target)
		}
	}
}
