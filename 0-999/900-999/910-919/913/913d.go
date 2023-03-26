// 00
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
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

type Pair struct {
	a, index, i int
	t           int64
}
type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].t > pq[j].t
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Pair)
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

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Pair, value int, priority int) {
// 	item.t = value
// 	item.t = priority
// 	heap.Fix(pq, item.index)
// }

func main() {
	var n int
	var T int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &T)
	pairs := make([]Pair, n)
	for i := 0; i < n; i++ {
		var a int
		var t int64
		fmt.Fscan(reader, &a, &t)
		pairs[i] = Pair{a: a, t: t, index: i, i: i + 1}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].a > pairs[j].a })
	pq := make(PriorityQueue, 0)
	time := int64(0)
	for i := 0; i < n; i++ {
		if time+pairs[i].t <= T {
			time += pairs[i].t
			heap.Push(&pq, &pairs[i])
		} else {
			if len(pq) > 0 {
				item := heap.Pop(&pq).(*Pair)
				if item.t > pairs[i].t {
					heap.Push(&pq, &pairs[i])
					time += pairs[i].t
					time -= item.t
				} else {
					heap.Push(&pq, item)
				}
			}
		}
		if len(pq) >= pairs[i].a {
			break
		}
	}
	mi := 1000000
	for i := 0; i < len(pq); i++ {
		mi = min(mi, pq[i].a)
	}
	mi = min(mi, len(pq))
	sort.Slice(pq, func(i, j int) bool { return pq[i].a < pq[j].a })
	ll := len(pq)
	for i := 0; i < len(pq); i++ {
		mi = max(mi, min(pq[i].a, ll-i))
	}
	//pq = pq[ll-mi:]
	write(f, mi, "\n")
	write(f, len(pq), "\n")
	for i := 0; i < len(pq); i++ {
		write(f, pq[i].i, " ")
	}
	write(f, "\n")

}
