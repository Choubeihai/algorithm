package main

func longestConsecutive(nums []int) int {
	var m = make(map[int]bool)
	var res int
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for num := range m {
		if m[num-1] {
			continue
		} else {
			var tmp = 0
			for m[num] {
				tmp++
				num++
			}
			res = max(res, tmp)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
