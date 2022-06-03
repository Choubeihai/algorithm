package main

import "fmt"

func main() {
	arr := []float32{1, 2, 3, 4, 5, 6}
	fmt.Println(getMedium(arr))
}

func getMedium(arr []float32) float32 {
	n := len(arr)
	left := 0
	right := n - 1
	res := 0
	for {
		res = partition(arr, left, right)
		if res == n/2 {
			break
		} else if res > n/2 {
			right = res - 1
		} else {
			left = res + 1
		}
	}

	if n%2 == 0 {
		return (arr[res] + arr[res-1]) / 2
	} else {
		return arr[res]
	}
}

func partition(arr []float32, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			left++
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right--
		}
	}
	arr[left] = pivot
	return left
}
