package main

import "fmt"

func main() {
	b := 'a'
	fmt.Printf("%T ", string(b))
	fmt.Println(string(b))
}

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int) //m[value][index]
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if index, ok := m[value]; ok {
			return []int{i, index}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
