package main

func isNumber(s string) bool {
	var existE bool   // 表示已经存在E
	var existNum bool // 截止到目前index，是一个数字
	var isFloat bool  // 是小数
	for i := 0; i < len(s); i++ {
		var c = s[i]
		// 符号位只能在首位或E的后一位
		if (c == '+' || c == '-') && (s[i-1] == 'e' || s[i-1] == 'E' || i == 0) {
			continue
		} else if (c == 'e' || c == 'E') && !existE && existNum { //只存在一个E, 前面必须有数字, 后面也必须有数字
			existE = true
			existNum = false
		} else if c == '.' && !isFloat && !existE { //只存在一个小数点, 不能在E的后面

			isFloat = true
		} else if c-'0' >= 0 && c-'0' <= 9 {
			existNum = true
		} else {
			return false
		}
	}
	return existNum
}
