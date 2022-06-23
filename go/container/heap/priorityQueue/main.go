package main

import (
	"container/heap"
	"fmt"
)

// Entry 是 priorityQueue 中的元素
type Entry struct {
	// value 是 Entry 中的数据，可以是任意类型，这里使用 string
	value string

	// priority 定义 Entry 在 priorityQueue 中的优先级
	priority int

	// index 是 Entry 在 heap 中的索引位置
	// Entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 Entry.priority 一直不变的话，可以删除 index
	index int
}

// PriorityQueue 实现 heap.Interface 接口方法
type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 这里生成小根堆
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

// Push 往 priorityQueue 中放 Entry
func (pq *PriorityQueue) Push(x interface{}) {
	entry := x.(*Entry)
	entry.index = len(*pq)
	*pq = append(*pq, entry)
}

// Pop 从 priorityQueue 中取出最优先的 Entry
func (pq *PriorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	x.index = -1 // for safety
	*pq = (*pq)[:len(*pq)-1]
	return x
}

// update 更改 Entry 在 priorityQueue 的值和优先级
func (pq *PriorityQueue) update(entry *Entry, value string, priority int) {
	entry.value = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}

func main() {
	pq := &PriorityQueue{
		{value: "Alice", priority: 5},
		{value: "Bob", priority: 2},
		{value: "Allen", priority: 4},
		{value: "An", priority: 1},
		{value: "Angel", priority: 3},
	}

	// 初始化优先级队列 pq
	heap.Init(pq)
	// 往 pq 中添加 entry
	entry := &Entry{value: "Claudia", priority: 0}
	heap.Push(pq, entry)
	// 更改 entry 的优先级
	pq.update(entry, entry.value, 6)
	for pq.Len() != 0 {
		entry := heap.Pop(pq).(*Entry)
		// fmt.Printf("priority:%d, value:%s\n", entry.priority, entry.value)
		fmt.Println(entry.index, entry.priority, entry.value)
	}
}
