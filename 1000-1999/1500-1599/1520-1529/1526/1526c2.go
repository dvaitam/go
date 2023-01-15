// 00
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	//fmt.Println("before append", h)
	*h = append(*h, x.(int))
	//fmt.Println("after append", h)
}

func (h *IntHeap) Pop() any {
	old := *h
	//fmt.Println("before pop", h)
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	h := &IntHeap{}
	heap.Init(h)
	ans := 0
	health := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i] >= 0 || health+int64(a[i]) >= 0 {
			health += int64(a[i])
			ans++
			if a[i] < 0 {
				heap.Push(h, a[i])
			}
		} else {
			if h.Len() > 0 {
				if a[i] > (*h)[0] {
					health -= int64((*h)[0])
					heap.Pop(h)
					heap.Push(h, a[i])
					health += int64(a[i])
				}
			}
		}
	}
	write(f, ans, "\n")

}
