package slice

import "fmt"

/**
使用make内置函数创建切片时，第一个参数是类型，第二个参数是切片的初始长度，第三个参数是底层数组的长度，
当append时超出了底层数组的长度，会自动进行扩容，append是添加到切片的末尾。
*/
func createSlice1() {
	p := make([]int, 5, 9)
	fmt.Println(p) // [0,0,0,0,0]
}

// 切片(slice)分片和赋值的相关的问题
func createSlice2() {
	buffer := make([]int, 8)
	a := buffer[0:0]
	fmt.Println("a = ", a)        // a =  []
	fmt.Println("len = ", len(a)) // len =  0
	fmt.Println("cap = ", cap(a)) // cap =  8

	b := buffer[4:4]
	fmt.Println("b = ", b)        // b =  []
	fmt.Println("len = ", len(b)) // len = 0
	fmt.Println("cap = ", cap(b)) // cap = 4

	c := buffer[8:8]
	fmt.Println("c = ", c)        // c = []
	fmt.Println("len = ", len(c)) // len = 0
	fmt.Println("cap = ", cap(c)) // cap = 0
}
