package main

import (
	"fmt"
	"strings"
)

/**
根据中序表达式，求后序表达式（逆波兰表达式）

遍历中缀表达式中的每一个单词符号 x：
    如果 x 是一个操作数，直接将 x 追加到输出队列 Q 末尾，否则往下检查；
    如果 x 是一个左括号“(”，将 x 压入操作符栈 S 栈顶，否则往下检查；
    如果 x 是一个操作符：
        如果操作符栈 S 栈顶为一个优先级大于等于 x 的操作符，则将 S 栈顶的运算符弹出并且追加到输出队列 Q 末尾，最后将 x 压入栈顶；
        如果操作符栈 S 栈顶为一个优先级小于 x 的操作符，或者不为操作符（在这个简化算法里，只有可能是左括号），则直接将 x 压入栈顶即可。
    如果 x 是一个右括号“)”，则将操作符栈 S 栈顶到往下第一个左括号“(”之间的元素依次弹出并且追加到输出队列末尾，将“(”出栈丢弃，x 也不用入栈。注意：如果栈到底后仍没有找到左括号，则说明表达式不合法，左右括号不匹配。
最后将栈 S 中的元素全部依次弹出并且入队列 Q。
*/

func main() {
	s := "(2+5)*(8-1)-6*6"
	// 2 5 + 8 1 - * 6 6 * -
	fmt.Println(translate3(s))
}

func translate3(s string) string {
	var n = len(s)
	var stack = make([]byte, n) // 栈
	var p = -1
	var q []byte // 队列
	var res = &strings.Builder{}
	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			q = append(q, s[i])
		} else {
			if s[i] == '(' {
				p++
				stack[p] = s[i]
			} else if s[i] == ')' {
				for p >= 0 {
					if stack[p] == '(' {
						p--
						break
					} else {
						q = append(q, stack[p])
						p--
					}
				}
			} else if s[i] == '+' || s[i] == '-' {
				for p >= 0 {
					if stack[p] == '*' || stack[p] == '/' {
						q = append(q, stack[p])
						p--
					} else {
						break
					}
				}
				p++
				stack[p] = s[i]
			} else {
				p++
				stack[p] = s[i]
			}
		}
	}

	for p >= 0 {
		q = append(q, stack[p])
		p--
	}

	for i := 0; i < len(q); i++ {
		res.WriteByte(q[i])
		res.WriteByte(' ')
	}
	return res.String()
}
