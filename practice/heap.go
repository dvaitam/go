package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	fmt.Println("before append", h)
	*h = append(*h, x.(int))
	fmt.Println("after append", h)
}

func (h *IntHeap) Pop() any {
	old := *h
	fmt.Println("before pop", h)
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{10, 6, 4}
	heap.Init(h)
	//fmt.Println(h)
	heap.Push(h, 3)
	fmt.Println("after after", h)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Println("before before", h)
		fmt.Printf("%d \n", heap.Pop(h))
	}
}
