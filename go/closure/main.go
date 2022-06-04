package main

func main() {
	// 例子1
	x := 1
	// f就是一个闭包
	f := func() {
		println(x)
	}
	f() // 此时调用闭包，f会对x进行解引用，此时x的值为1，因此输出为1
	x = 2
	f() // 此时调用闭包，f会对x进行解引用，此时x的值为2，因此输出为2
	x = 3
	f() // 此时调用闭包，f会对x进行解引用，此时x的值为3，因此输出为3

	// 例子2
	var arr []func()
	for i := 0; i < 3; i++ {
		arr = append(arr, func() {
			println(i)
		})

	}
	for j := 0; j < 3; j++ {
		arr[j]() // 输出为3，3，3，这是因为arr里面的三个闭包里保存的是i的地址，执行闭包时才会对i进行解引用，
		// i按理说已经销毁，但是遇到闭包发生了内存逃逸
	}

	// 例子3
	var arr2 []func()
	for i := 0; i < 3; i++ {
		j := i
		arr = append(arr2, func() {
			println(j)
		})

	}
	for i := 0; i < 3; i++ {
		arr2[i]() // 输出结果：0，1，2。就不需要解释了吧
	}
}
