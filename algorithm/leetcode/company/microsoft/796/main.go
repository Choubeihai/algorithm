package main

func rotateString(s string, goal string) bool {
	var n = len(s)
	var m = len(goal)
	if n != m {
		return false
	}
	for i := 0; i < n; i++ {
		var flag = true
		for j := 0; j < n; j++ {
			if s[(j+i)%n] != goal[j] {
				flag = false
				break
			}
		}
		if flag == true {
			return true
		}
	}
	return false

}
