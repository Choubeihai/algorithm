package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			c <- 1
		}
	}()

	for {
		select {
		case <-c:
			fmt.Println("111111")
		case <-time.After(2 * time.Second):
			fmt.Println("2222222222")
		}
	}

}
