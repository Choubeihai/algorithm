package main

import "fmt"

/**
一个整数数组，长度为n，将其分为m份，使各份的和相等，求m的最大值
    比如{3，2，4，3，6} 可以分成{3，2，4，3，6} m=1;
    {3,6}{2,4,3} m=2
    {3,3}{2,4}{6} m=3 所以m的最大值为3
*/

func main() {
	//var arr = []int{3, 2, 4, 3, 6}
	//var arr = []int{3, 3, 2, 2}
	var arr = []int{3, 1, 5, 1}
	fmt.Println(divideArray(arr))

}

func divideArray(arr []int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	var m = n
	for m > 0 {
		tag := make([]bool, n)
		if sum%m != 0 {
			m--
		} else {
			subSum := sum / m
			i := 0
			for ; i < n; i++ {
				if tag[i] {
					continue
				}
				if !recur(arr, tag, subSum) {
					break
				}
			}

			if i == n {
				break
			}
			m--
		}
	}
	return m
}

func recur(arr []int, tag []bool, sum int) bool {
	if sum == 0 {
		return true
	}

	for i := 0; i < len(arr); i++ {
		if tag[i] {
			continue
		}
		if sum-arr[i] < 0 {
			continue
		}
		tag[i] = true
		flag := recur(arr, tag, sum-arr[i])
		if flag {
			return true
		} else {
			tag[i] = false
		}

	}

	return false
}
