package main

import "sync"

type singleton struct {
}

// sync.Once
var instance *singleton
var once sync.Once

func getInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
		})
	}
	return instance
}
