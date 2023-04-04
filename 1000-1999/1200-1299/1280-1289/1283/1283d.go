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

type Item struct {
	index    int
	priority int
	vertex   int
	left     bool
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	ans := int64(0)
	x := make([]int, n)
	occupied := map[int]bool{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		occupied[x[i]] = true
	}
	// sort.Ints(x)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for i := 0; i < n; i++ {
		if !occupied[x[i]-1] {
			heap.Push(&pq, &Item{priority: 0, vertex: x[i], left: true})
		}
		if !occupied[x[i]+1] {
			heap.Push(&pq, &Item{priority: 0, vertex: x[i]})
		}
	}
	y := make([]int, 0)
	for i := 0; i < m; i++ {
		for true {
			item := heap.Pop(&pq).(*Item)
			if item.left {
				index := item.vertex - item.priority - 1
				if !occupied[index] {
					ans += int64(item.priority + 1)
					y = append(y, index)
					occupied[index] = true
					item.priority++
					heap.Push(&pq, item)
					break
				}
			} else {
				index := item.vertex + item.priority + 1
				if !occupied[index] {
					ans += int64(item.priority + 1)
					y = append(y, index)
					occupied[index] = true
					item.priority++
					heap.Push(&pq, item)
					break
				}
			}
		}

	}
	write(f, ans, "\n")
	for i := 0; i < m; i++ {
		write(f, y[i], " ")
	}
	write(f, "\n")

}
