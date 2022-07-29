package main

import "fmt"

func main() {

}

func f1() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for i := 0; i < 10; i++ {
		data, ok := <-c
		fmt.Printf("data=%v, isOpen=%v \n", data, ok)
	}
}
