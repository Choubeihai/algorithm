package main

import "fmt"

/**
编译时动态注入变量信息
*/

var (
	version   string
	BuildTime string
	GoVersion string
)

func main() {
	fmt.Printf("%s\n%s\n%s\n", version, BuildTime, GoVersion)
}

// go build -ldflags "-X main.version=1.0.0 -X 'main.BuildTime=`date`' -X 'main.GoVersion=`go version`'" -o main main.go
