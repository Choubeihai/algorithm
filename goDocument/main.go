package main

import "sync"

func main() {
	mutex := sync.Mutex{}
	defer mutex.Unlock()
	mutex.Lock()
}
