package main

func convert(s string, numRows int) string {
	n := len(s)
	r := numRows          // 行数
	if r == 1 || r >= n { // 当行数为1，直接返回，当行数大于字符串长度，也直接返回
		return s
	}
	t := r + r - 2          // 一个周期占用的字符数
	numT := (n + t - 1) / t // 一共有多少个周期
	cc := 1 + r - 2         // 一个周期占用的列数
	c := cc * numT          // 列数

	matrix := make([][]byte, r)
	for i := 0; i < r; i++ {
		matrix[i] = make([]byte, c)
	}

	x := 0
	y := 0
	for i := 0; i < len(s); i++ {
		matrix[x][y] = s[i]
		if i%t < r-1 { // 这个是关键的判断，如果当前i%t < r-1，说明当前i不在最后一行，可以往下移动
			x++ // 向下移动
		} else {
			x-- // 向右上移动
			y++
		}
	}
	res := make([]byte, 0, n)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] > 0 {
				res = append(res, matrix[i][j])
			}
		}
	}
	return string(res)
}
