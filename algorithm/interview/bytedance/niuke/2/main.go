package main

import (
	"fmt"
	"sort"
)

/**
一堆区间， 例如 [1,2],[3,4][2,9] 等，求这堆区间覆盖在数轴上的长度，区间无序
*/

func main() {
	var arr = [][]int{{1, 2}, {3, 4}, {2, 9}}
	ans := findArrLength(arr)
	fmt.Println(ans)

}

func findArrLength(arr [][]int) int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[i][1]
		}
		return arr[i][0] < arr[j][0]
	})
	var res [][]int
	res = append(res, arr[0])
	for i := 1; i < len(arr); i++ {
		index := len(res) - 1
		if arr[index][1] > arr[i][1] {
			continue
		} else if arr[index][1] > arr[i][0] {
			tmp := []int{res[index][0], arr[i][1]}
			res[index] = tmp
		} else {
			res = append(res, arr[i])
		}
	}
	var ans int
	for i := 0; i < len(res); i++ {
		ans += res[i][1] - res[i][0]
	}
	return ans
}
