package main

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	var res [][]int
	added := false //是否已经被添加进去

	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < left { // 新成员在左边且无交叉
			res = append(res, intervals[i])
		} else if intervals[i][0] > right { // 新成员在右边且无交叉
			if !added {
				res = append(res, []int{left, right})
				added = true
			}
			res = append(res, intervals[i])
		} else { // 有交叉，直接min  max 判断
			left = min(left, intervals[i][0])
			right = max(right, intervals[i][1])
		}
	}

	if !added {
		res = append(res, []int{left, right})
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
