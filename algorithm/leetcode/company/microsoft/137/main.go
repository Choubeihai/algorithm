package main

func singleNumber(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		var count int32 = 0
		for j := 0; j < len(nums); j++ {
			num := int32(nums[j])
			count = count + ((num >> i) & 1)
		}
		if count%3 != 0 {
			res = res | (1 << i)
		}
	}
	return int(res)
}
