package main

/**
作者：Libre
链接：https://www.nowcoder.com/discuss/615675?type=post&order=recall&pos=&page=1&ncTraceId=&channel=-1&source_id=search_post_nctrack&gio_id=7D56EE02BFBA6CAD01FFECB2B6EE4D8D-1659257344742
来源：牛客网
编程题，正整数数组，数字代表可以走几步，问能不能走到最后，例如
[1,2,3,4], 可以从位置1 走1步到2， 从位置2走2步到4，能走到最后，[1,0,3,4,5],位置1只能走1步，走到位置2不能再走
*/

func canReach(arr []int) bool {
	reach := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		reach[i] = i + arr[i]
	}
	var canMaxReach = reach[0]
	for i := 1; i < len(arr); i++ {
		if canMaxReach >= i {
			canMaxReach = max(canMaxReach, reach[i])
		} else {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
