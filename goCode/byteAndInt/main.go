package main

import "fmt"

func main() {
	var a byte = '1'
	fmt.Println(a)       // 输出49
	fmt.Println(int(a))  // 输出49
	fmt.Println(a - '0') // 输出1

	var b int = 49
	fmt.Println(b)        // 输出49
	fmt.Println(byte(b))  // 输出49
	fmt.Printf("%q\n", b) // 输出'1'
	fmt.Printf("%c\n", b) // 输出 1
}
