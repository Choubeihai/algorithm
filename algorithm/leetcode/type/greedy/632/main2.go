package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	var nums = [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	fmt.Println(smallestRange(nums))
}

func smallestRange(nums [][]int) []int {
	numsC = nums
	next = make([]int, len(numsC))
	h := new(MyHeap)
	heap.Init(h)

	left := math.MinInt / 2
	right := math.MaxInt / 2
	minRange := right - left
	nowMax := math.MinInt
	for i := 0; i < len(numsC); i++ {
		if numsC[i][0] > nowMax {
			nowMax = numsC[i][0]
		}
		heap.Push(h, i)
	}
	for {
		index := heap.Pop(h).(int)
		nowRange := nowMax - numsC[index][next[index]]
		if nowRange < minRange {
			minRange = nowRange
			left = numsC[index][next[index]]
			right = nowMax
		}
		next[index]++
		if next[index] == len(numsC[index]) {
			break
		}
		heap.Push(h, index)
		if nowMax < numsC[index][next[index]] {
			nowMax = numsC[index][next[index]]
		}
	}
	return []int{left, right}
}

type MyHeap []int //存储下标

var (
	numsC [][]int // 作为数组s，第一个数组、第二个数组....
	next  []int   // 作为数组的下标，例如next[4]=2，表示第4个数组的下标目前为2
)

func (heap MyHeap) Len() int {
	return len(heap)
}

func (heap MyHeap) Less(i, j int) bool {
	return numsC[heap[i]][next[heap[i]]] < numsC[heap[j]][next[heap[j]]]
}

func (heap MyHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MyHeap) Push(x interface{}) {
	*heap = append(*heap, x.(int))
}

func (heap *MyHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[0 : n-1]
	return x
}
