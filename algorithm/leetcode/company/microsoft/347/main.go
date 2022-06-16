package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1}
	k := 1
	fmt.Println(topKFrequent2(nums, k))
}

// 基于快速排序
func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	kv := make([][2]int, 0, len(m))
	for key, value := range m {
		kv = append(kv, [2]int{key, value})
	}
	res := find(kv, 0, len(kv)-1, k-1)
	return res
}

func find(kv [][2]int, left, right, k int) []int {
	if left <= right {
		p := partition(kv, left, right)
		if p == k {
			var res []int
			for i := 0; i <= k; i++ {
				res = append(res, kv[i][0])
			}
			return res

		} else if p < k {
			return find(kv, p+1, right, k)
		} else {
			return find(kv, left, p-1, k)
		}
	} else {
		return nil
	}
}

func partition(nums [][2]int, left, right int) int {
	i := left
	j := right
	pivot := nums[i]
	for i < j {
		for i < j && nums[j][1] <= pivot[1] {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i][1] > pivot[1] {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = pivot
	return i
}

// 基于堆
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
