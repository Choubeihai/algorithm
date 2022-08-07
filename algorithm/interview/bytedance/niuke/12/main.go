package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	var res [][]int
	res = append(res, intervals[0])
	fmt.Println(res)
	for i := 1; i < len(intervals); i++ {
		var n = len(res)
		if intervals[i][0] > res[n-1][1] {
			res = append(res, intervals[i])
		} else {
			if intervals[i][1] > res[n-1][1] {
				res[n-1] = []int{res[n-1][0], intervals[i][1]}
			}
		}
	}
	return res
}
