package main

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	var res int

	for i := 0; i < n; i++ {
		m := make(map[int]int)
		for j := i + 1; j < n; j++ {
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			m[x*100000+y]++
		}
		for _, count := range m {
			res = max(res, count+1)
		}
	}
	return res
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
