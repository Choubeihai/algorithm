package main

func main() {

}

// 方法一
func binarySearch1(arr []int, left, right int, value int) int {
	for left < right {
		mid := (left + right) / 2
		if arr[mid] < value {
			left = mid + 1
		} else if arr[mid] > value {
			right = mid
		} else {
			return arr[mid]
		}
	}
	return -1
}

// 方法二
func binarySearch2(arr []int, left, right int, value int) int {
	for left < right {
		mid := (left + right + 1) / 2
		if arr[mid] < value {
			left = mid
		} else if arr[mid] > value {
			right = mid - 1
		} else {
			return arr[mid]
		}
	}
	return -1
}
