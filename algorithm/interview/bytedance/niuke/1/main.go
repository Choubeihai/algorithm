package main

import "fmt"

/**
一道字节算法题

链接：算法连接
来源：牛客网
把数组元素按照正负序重排列
给定一个数组，数组它按照下面的规则重排列后的数组： 1. 数组中的正负数相互间隔 2. 符号相同的数字相对顺序不变 3. 如果某种符号的数字多余，放到数组最后
例如：-1,3,2,4,5,-6,7,-9
重排列后：3,-1,2,-6,4,-9,5,7
空间复杂度要求O(1)
*/

func main() {
	var arr = []int{-1, 3, 2, 4, 5, -6, 7, -9}
	sortArr(arr)
	fmt.Println(arr)
}

func sortArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			if arr[i] > 0 {
				continue
			}
			var j = i + 1
			for ; j < len(arr); j++ {
				if arr[j] > 0 {
					break
				}
			}
			if j == len(arr) {
				break
			}
			var v = arr[j]
			for k := j; k > i; k-- {
				arr[k] = arr[k-1]
			}
			arr[i] = v
		} else {
			if arr[i] < 0 {
				continue
			}
			var j = i + 1
			for ; j < len(arr); j++ {
				if arr[j] < 0 {
					break
				}
			}
			if j == len(arr) {
				break
			}
			var v = arr[j]
			for k := j; k > i; k-- {
				arr[k] = arr[k-1]
			}
			arr[i] = v
		}
	}
}
