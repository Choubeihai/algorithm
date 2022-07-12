package s1

import "sync"

type singleton struct {
}

// 双重检验加锁, 懒汉式
var instance *singleton
var mu sync.Mutex

func getInstance() *singleton {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			instance = &singleton{}
		}
		mu.Unlock()
	}
	return instance
}
