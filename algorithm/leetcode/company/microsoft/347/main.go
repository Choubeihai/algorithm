package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	h := &H{}
	heap.Init(h)
	for key, value := range m {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		pop := heap.Pop(h).([2]int)
		res[i] = pop[0]
	}
	return res
}

type H [][2]int

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *H) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
