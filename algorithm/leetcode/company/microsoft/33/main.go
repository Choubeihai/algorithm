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
	var n = len(nums)
	return find(nums, 0, n-1, target)
}

func find(nums []int, left, right int, target int) int {
	/**
	不需要考虑left > right的情况，因为mid=(left+right)/2，左侧是[left, mid], 右侧是[mid+1, right]
	mid计算公式决定了mid一定小于right， 可能等于left，因此不会越界。
	*/
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	mid := (left + right) / 2
	if nums[left] <= nums[mid] { // 左侧有序，一定要加=，因为left可能等于mid
		if target >= nums[left] && target <= nums[mid] {
			return find(nums, left, mid, target)
		} else {
			return find(nums, mid+1, right, target)
		}
	} else { // 右侧有序
		if target >= nums[mid+1] && target <= nums[right] {
			return find(nums, mid+1, right, target)
		} else {
			return find(nums, left, mid, target)
		}
	}
}
