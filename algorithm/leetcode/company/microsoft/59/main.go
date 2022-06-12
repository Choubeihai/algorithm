package main

import "fmt"

func main() {
	n := 2
	fmt.Println(generateMatrix(n))

}

func generateMatrix(n int) [][]int {
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	var top = 0
	var bottom = n - 1
	var left = 0
	var right = n - 1
	var value = 1

	for top <= bottom && left <= right {
		// 左 ——> 右
		for i := left; i <= right; i++ {
			res[top][i] = value
			value++
		}
		// 上 ——> 下
		for i := top + 1; i <= bottom; i++ {
			res[i][right] = value
			value++
		}

		if top < bottom && left < right {
			// 右——>左
			for i := right - 1; i >= left; i-- {
				res[bottom][i] = value
				value++
			}
			for i := bottom - 1; i > top; i-- {
				res[i][left] = value
				value++
			}
		}
		top++
		bottom--
		right--
		left++
	}
	return res
}
