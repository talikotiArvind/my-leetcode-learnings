package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))              // 5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))     // 4
}