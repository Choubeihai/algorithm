package reentryMutex

import (
	"fmt"
	goid2 "github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

type ReentryMutex struct {
	mutex   sync.Mutex // 锁
	owner   int64      // go程id
	reentry int32      // 拿到owner以后才能去拿到reentry
}

func (m *ReentryMutex) Lock() {
	goid := goid2.Get()
	if goid == atomic.LoadInt64(&m.owner) {
		m.reentry++
		return
	}
	m.mutex.Lock()
	atomic.StoreInt64(&m.owner, goid)
	m.reentry = 1
}

func (m *ReentryMutex) Unlock() {
	goid := goid2.Get()
	if atomic.LoadInt64(&m.owner) != goid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, goid))
	} else {
		m.reentry--
		if m.reentry == 0 {
			atomic.StoreInt64(&m.owner, -1)
			m.mutex.Unlock()
		}
	}
}
