package main

func singleNumber(nums []int) []int {
	var xor int
	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
	}
	var bit = 1
	for xor&bit == 0 {
		bit = bit << 1
	}
	var first, second int
	for i := 0; i < len(nums); i++ {
		if nums[i]&bit == 0 {
			first = first ^ nums[i]
		} else {
			second = second ^ nums[i]
		}
	}
	return []int{first, second}
}
