# sync.RWMutex

## 实现读写锁的目的

1. 场景：读多写少
2. 写锁需要同时阻塞写锁和读锁：一个协程拥有写锁时，其他协程的无论【读锁】还是【写锁】都要阻塞；
3. 读锁需要阻塞写锁：一个协程拥有读锁时，其他协程的【写锁】需要阻塞；
4. 读锁不阻塞读锁：一个协程拥有读锁是时，其他协程也可以拥有【读锁】。

```go
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // 用于writer等待 读完成 排队的信号量
	readerSem   uint32 // 用于reader等待 写完成 排队的信号量
	readerCount int32  // 读锁的计数器
	readerWait  int32  // 等待读锁释放的数量
}
```

最多可以有1<<30个读并发协程，即2的30此次。


## 实现

读锁其实就是无锁，写锁基于Mutex
1. 读锁加锁：会首先判断readCount+1后是否小于0,如果小于，则说明正在写，直接被阻塞；如果大于，则直接读；
2. 读锁解锁，readCount--,如果readCount==0而且当前有阻塞的写协程，那么唤醒被阻塞的写协程；否则直接退出；
3. 写锁加锁：写锁执行Lock函数，将readCount的值复制到readWait中，用来标记排在写操作前面的协程，当排在前面的Reader完事以后，同时递减ReadCount和ReadWait，当ReadWait为0时，当readCount为0时，唤醒写操作，将ReadCount减去1<<30；
4. 写操作解锁：写操作执行readCount+1<<30,然后执行UnLock。
