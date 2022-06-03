package main

func main() {
	s := "xywrrmp"
	t := "xywrrmu#p"
	println(backspaceCompare(s, t))
}

func backspaceCompare(s string, t string) bool {
	stackS := make([]byte, len(s))
	stackT := make([]byte, len(t))
	sp := -1 // 栈顶所在位置
	tp := -1 // 栈顶所在位置
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if sp >= 0 {
				sp--
			}
		} else {
			sp++
			stackS[sp] = s[i]
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if tp >= 0 {
				tp--
			}
		} else {
			tp++
			stackT[tp] = t[i]
		}
	}
	for sp >= 0 && tp >= 0 {
		if stackS[sp] != stackT[tp] {
			return false
		} else {
			sp--
			tp--
		}
	}
	if sp != -1 {
		return false
	} else if tp != -1 {
		return false
	} else {
		return true
	}
}
