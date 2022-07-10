package main

import (
	"container/heap"
	"sync"
	"sync/atomic"
	"time"
)

type DelayQueue struct {
	C chan interface{} // 接收过期元素

	mu sync.Mutex
	pq priorityQueue

	// Similar to the sleeping state of runtime.timers.
	sleeping int32 // 表示 delay queue中没有任何元素，所以delay queue处于休眠状态
	wakeupC  chan struct{}
}

// New creates an instance of delayQueue with the specified size.
func New(size int) *DelayQueue {
	return &DelayQueue{
		C:       make(chan interface{}),
		pq:      newPriorityQueue(size),
		wakeupC: make(chan struct{}),
	}
}

// Offer inserts the element into the current queue.
func (dq *DelayQueue) Offer(elem interface{}, expiration int64) {
	item := &item{
		Value:    elem,
		Priority: expiration,
	}

	dq.mu.Lock()
	heap.Push(&dq.pq, item)
	index := item.Index
	dq.mu.Unlock()

	if index == 0 {
		// A new item with the earliest expiration is added.
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeupC <- struct{}{}
		}
	}
}

// Poll 开启一个无限循环，在循环中持续等待一个元素过期，然后将这个过期元素发送到C中。
func (dq *DelayQueue) Poll(nowF func() int64) {
	for {
		now := nowF()
		dq.mu.Lock()
		item, delta := dq.pq.PeekAndShift(now)
		// 说明delay queue 中没有元素，或者头部元素没有过期
		if item == nil {
			// 直接睡觉
			atomic.StoreInt32(&dq.sleeping, 1)
		}
		dq.mu.Unlock()

		if item == nil {
			if delta == 0 { // 说明没有元素
				select {
				case <-dq.wakeupC:
					continue
				}
			} else if delta > 0 { //说明有元素，但是没有到过期时间
				select {
				case <-dq.wakeupC:
					continue
				case <-time.After(time.Duration(delta) * time.Millisecond):
					if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
						<-dq.wakeupC
					}
					continue
				}
			}
		}
		select {
		case dq.C <- item.Value:
			// The expired element has been sent out successfully.
		}
	}

}
