package main

import "fmt"

/**
  堆排序 非稳定算法
*/

func main() {
	var arr = []int{2, 1, 3, 1, 3, 5, 9}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}

// 大顶堆排序
func heapSort(arr []int) {
	// 构建大顶堆,过程：保证父节点的值永远比子节点的值大，但是此时无法保证兄弟节点之间的大小关系
	len := len(arr)
	for i := len/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len)
	}
	// 交换堆顶元素，过程：堆排序保证了父节点的值比子节点都大，所以根节点是最大值，每一次循环都是根节点和最后的节点对换，然后在顺势调整堆节点，
	// 第二大的节点就变成了根节点，然后在替换根节点和最后的一个节点....
	for i := len - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjustHeap(arr, 0, i)
	}
}

// adjustHeap 调整parentIndex下的所有子节点的位置，保证父节点的值是最大的，兄弟节点之间无要求
func adjustHeap(arr []int, parentIndex int, len int) {
	val := arr[parentIndex]
	childIndex := 2*parentIndex + 1
	for childIndex < len {
		if childIndex+1 < len && arr[childIndex] < arr[childIndex+1] {
			childIndex++
		}
		if arr[childIndex] > val {
			arr[parentIndex] = arr[childIndex]
			parentIndex = childIndex
		} else {
			break
		}
		childIndex = childIndex*2 + 1
	}
	arr[parentIndex] = val
}
