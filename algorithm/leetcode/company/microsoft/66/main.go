package main

func plusOne(digits []int) []int {
	var bit = 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := bit + digits[i]
		digits[i] = sum % 10
		bit = sum / 10
	}
	if bit != 0 {
		var res = []int{1}
		res = append(res, digits...)
		return res

	} else {
		return digits
	}
}
