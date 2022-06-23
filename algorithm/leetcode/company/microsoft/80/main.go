package main

import "fmt"

func main() {
	var nums = []int{1, 1, 1, 2, 2, 3}
	n := removeDuplicates(nums)
	fmt.Println(n)
}

/**
前提：是排好序的数组
思路：快慢指针slow，fast。slow指向的是下一个将要被填充的位置，fast指针指向的是需要比较的位置。
     如果nums[fast] == nums[slow-2],说明前面已经有两个相同的字符，那么fast位置的字符不需要
     再要了,直接移动。如果nums[fast] ！= nums[slow-2]，说明此时可以填充，令nums[slow] = nums[fast]，
     slow和fast同时移动

*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	var slow = 2
	var fast = 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
