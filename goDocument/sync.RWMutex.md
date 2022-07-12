# sync.RWMutex

## 实现读写锁的目的

1. 场景：读多写少
2. 写锁需要同时阻塞写锁和读锁：一个协程拥有写锁时，其他协程的无论【读锁】还是【写锁】都要阻塞；
3. 读锁需要阻塞写锁：一个协程拥有读锁时，其他协程的【写锁】需要阻塞；
4. 读锁不阻塞读锁：一个协程拥有读锁是时，其他协程也可以拥有【读锁】。

```go
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```

读写锁支持的最大并发读协程的数量是 1 << 30，也就是 2 的 30次方个，这已经非常大了，几乎不会有日常的业务开发能触及到这个量级。

这里可以看出 RWMutex 也是基于 Mutex 的能力进行的包装，我们记得 Mutex 只包含了一个 state 代表状态，一个 sem 代表信号量。这里读写锁为了进一步拆分粒度。增加了下面几个变量：

- writerSem: 写阻塞等待的信号量，最后一个读者释放锁时，会释放该信号量；
- readerSem: 读阻塞的协程等待的信号量，持有写锁的协程释放锁后，会释放该信号量；
- readerCount: 拿到锁的 readers 数量
- readerWait: 正在写阻塞时的的 readers 数量。

## 实现

1. 写操作相互阻塞：写操作之间的控制依赖于内置的 Mutex
2. 写操作阻塞读操作：写操作是将 readerCount 变成负值来阻止读操作的。
3. 在执行写锁定，即 `Lock()` 时，会先将 readerCount 减去 1 << 30，这样就小于0了。再有读锁定到来的时候，检测 readerCount，发现为负数，就知道此时已经有写操作进行了，读需要阻塞。而真实的【读操作的数量】，只需要 readerCount 再加上 1 << 30 就拿到了。
4. **读操作阻塞写操作** ：写操作只要发现有正在进行的读操作，就应该停下来阻塞到这里。这个直接通过 readerCount 就可以判断。
5. **如何保证写操作不会饿死**：写操作上锁的 `Lock()` 函数中，会把 readerCount 的值复制到 readerWait，用来标记【排在写操作前面的 reader 个数】。当排在前面的 reader 完事的时候，会递减 readerCount 的值，并同时递减 readerWait 的值。当 readerWait 等于 0 时会唤醒写操作。一句话：写操作相当于把一段连续的读操作分成两个部分，前一部分完成的时候，通过 writerSem 唤醒写操作，实现插队的效果。