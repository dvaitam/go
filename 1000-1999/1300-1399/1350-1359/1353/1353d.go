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

// An Item is something we manage in a priority queue.
type Item struct {
	l        int
	r        int
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority == pq[j].priority {
		return pq[i].l < pq[j].l
	} else {
		return pq[i].priority > pq[j].priority
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func print_heap(f *bufio.Writer, pq *PriorityQueue) {
	for i := 0; i < len(*pq); i++ {
		write(f, (*pq)[i], " ")
	}
	write(f, "\n")
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n+1)
		h := make(PriorityQueue, 0)
		heap.Init(&h)
		heap.Push(&h, &Item{l: 1, r: n, priority: n})
		curr := 1
		for h.Len() > 0 {
			segment := heap.Pop(&h).(*Item)
			mid := (segment.l + segment.r) / 2
			//write(f, "assigning segement", segment, "mid is ", mid, "\n")

			a[mid] = curr
			curr++
			if mid < segment.r {
				//	write(f, "O am here", mid+1, segment.r, "\n")
				heap.Push(&h, &Item{l: mid + 1, r: segment.r, priority: segment.r - mid})
			}
			if segment.l < mid {
				//	write(f, "O am here", segment.l, mid-1, "\n")
				heap.Push(&h, &Item{l: segment.l, r: mid - 1, priority: mid - segment.l})
			}
			//print_heap(f, &h)

			// if curr > n {
			// 	break
			// }
			//	write(f, "after one ", h, "\n")
		}
		for i := 1; i <= n; i++ {
			write(f, a[i])
			if i == n {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	}
}
