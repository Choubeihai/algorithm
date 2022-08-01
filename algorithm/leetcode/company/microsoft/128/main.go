package main

func longestConsecutive(nums []int) int {
	var m = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	var res int
	for num := range m {
		if !m[num-1] {
			tmpRes := 0
			for m[num] {
				tmpRes++
				num++
			}
			if tmpRes > res {
				res = tmpRes
			}
		}
	}
	return res
}
