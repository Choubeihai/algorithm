package main

/**

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	go func(i int) {
		for {
			select {
			case <-ch1:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				ch2 <- 1
			}
		}
	}(1)
	letter := 'A'
	for {
		select {
		case <-ch2:
			if letter > 'Z' {
				return
			}
			fmt.Print(string(letter))
			letter++
			fmt.Print(string(letter))
			letter++
			ch1 <- 1
		}
	}

}
