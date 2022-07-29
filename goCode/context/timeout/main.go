package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

/**
超时控制，可能会有内存泄露的风险
*/

func job(ctx context.Context) error {
	context, _ := context.WithTimeout(ctx, time.Second*1)
	done := make(chan struct{}) // 不设置成有缓冲区的话，会有内存泄露的风险
	go func() {
		time.Sleep(1200 * time.Millisecond) // 模拟工作执行时间
		done <- struct{}{}                  // 如果done设置成无缓冲，没有协程消费done里面的数据，会导致该协程阻塞在这里。
	}()
	select {
	case <-done:
		return nil
	case <-context.Done():
		return errors.New("job time out")
	}
}

func main() {
	for i := 0; i < 500; i++ {
		ctx := context.Background()
		go func() {
			job(ctx)
		}()
	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(runtime.NumGoroutine())
	}
}
