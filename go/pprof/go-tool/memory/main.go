package main

import (
	"bytes"
	"math/rand"
	"os"
	"runtime/pprof"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// go:noinline
func generate(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(Letters[rand.Intn(len(Letters))])
	}
	return buf.String()
}

//go:noinline
func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}

	return result
}

func main() {
	// memory profile
	f, err := os.OpenFile("./memory-profile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	//err = pprof.WriteHeapProfile(f)
	//if err != nil {
	//	return
	//}

	for i := 0; i < 1000; i++ {
		repeat(generate(100), 100)
	}

	pprof.Lookup("heap").WriteTo(f, 0)

}
