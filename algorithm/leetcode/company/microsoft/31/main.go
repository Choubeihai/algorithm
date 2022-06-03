package main

import "fmt"

/**
1. 先找出最大的索引 k 满足 nums[k] < nums[k+1];如果不存在，就翻转整个数组,结束；
2. 再找出另一个最大索引 l 满足 nums[l] > nums[k]；
3. 交换 nums[l] 和 nums[k]；
4. 最后翻转 nums[k+1:]。

比如 nums = [1,2,7,4,3,1]，下一个排列是什么？
我们找到第一个最大索引是 nums[1] = 2
再找到第二个最大索引是 nums[4] = 3
交换，nums = [1,3,7,4,2,1];
翻转，nums = [1,3,1,2,4,7]

本质思路：
1.先倒序遍历数组, 找到第一个 (前一个数比后一个数小的位置) (即nums[i] < nums[i+1]);
2. 这个时候我们不能直接把后一个数nums[i+1] 跟前一个数nums[i]交换就完事了; 还应该从nums[i+1]-->数组末尾这一段的数据中
   找出最优的那个值( 如何最优? 即比nums[i]稍微大那么一丢丢的数, 也就是 nums[i+1]-->数组末尾中, 比nums[i]大的数中最小的那个值)
3. 找到之后, 跟num[i]交换, 这还不算是下一个排列, num[i]后面的数值还不够小, 所以还应当进升序排列
*/

func main() {
	var arr = []int{9, 8, 1, 9, 4, 6, 5, 5}
	nextPermutation(arr)
	fmt.Println(arr)
}

func nextPermutation(nums []int) {
	k := -1
	l := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}

	// 反转数组
	if k == -1 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return
	}

	for i := k + 1; i < len(nums); i++ {
		if nums[i] > nums[k] {
			l = i
		}
	}

	nums[k], nums[l] = nums[l], nums[k]
	for i, j := k+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}
