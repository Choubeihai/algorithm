package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {

	// cpu profile
	f, err := os.OpenFile("./cpu-profile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	err = pprof.StartCPUProfile(f)
	if err != nil {
		return
	}
	defer pprof.StopCPUProfile()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
