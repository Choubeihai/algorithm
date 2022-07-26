package main

func smallestK(arr []int, k int) []int {
	return quickSort(arr, 0, len(arr)-1, k-1)
}

func quickSort(arr []int, left, right int, k int) []int {
	if left > right {
		return nil
	}
	part := partition(arr, left, right)
	if part == k {
		var res []int
		for i := 0; i <= k; i++ {
			res = append(res, arr[i])
		}
		return res
	} else if part < k {
		return quickSort(arr, part+1, right, k)
	} else {
		return quickSort(arr, left, part-1, k)
	}
}

func partition(arr []int, left, right int) int {
	var i = left
	var j = right
	var pivot = arr[i]
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] < pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = pivot
	return i
}
