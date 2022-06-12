package main

import "sort"

func merge(intervals [][]int) [][]int {
	var n = len(intervals)
	if n == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < n; i++ {
		if right >= intervals[i][0] {
			if left >= intervals[i][0] {
				left = intervals[i][0]
			}
			if right <= intervals[i][1] {
				right = intervals[i][1]
			}

		} else { // 说明没有连在一起
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}

	res = append(res, []int{left, right})
	return res
}
