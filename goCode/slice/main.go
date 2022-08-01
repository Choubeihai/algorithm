package main

import "fmt"

func main() {
	f2()
}

/**
使用make内置函数创建切片时，第一个参数是类型，第二个参数是切片的初始长度，第三个参数是底层数组的长度，
当append时超出了底层数组的长度，会自动进行扩容，append是添加到切片的末尾。
*/
func createSlice1() {
	f2()
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

func f2() {
	// 切片遍历
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	// for循环方式
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v", i, slice[i])
	}
	fmt.Println()
	// for range方式
	for i := range slice {
		fmt.Println(i)
	}
}
