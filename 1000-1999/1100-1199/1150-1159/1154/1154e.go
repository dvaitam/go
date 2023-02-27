// 00
package main

import (
	"bufio"
	"container/heap"
	"container/list"
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
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
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
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
func give_max(pq *PriorityQueue, deleted map[int]bool) int {
	item := heap.Pop(pq).(*Item)
	for deleted[item.value] {
		item = heap.Pop(pq).(*Item)
	}
	return item.value

}
func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	ll := list.New()
	a := make([]int, n)
	m := map[int]*list.Element{}
	pq := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)
		a[i] = tmp
		ll.PushBack(tmp)
		m[tmp] = ll.Back()
		pq[i] = &Item{
			value:    tmp,
			priority: tmp,
			index:    i,
		}
	}
	heap.Init(&pq)

	first := "1"
	ans := map[int]string{}
	deleted := map[int]bool{}
	for ll.Len() > 0 {
		mx := give_max(&pq, deleted)

		//write(f, "max value is ", mx, "\n")
		deleted[mx] = true
		start := k
		start_item := m[mx]
		for start_item.Prev() != nil && start > 0 {
			int_r := start_item.Prev().Value.(int)
			ans[int_r] = first
			deleted[int_r] = true

			ll.Remove(start_item.Prev())
			start--
		}
		start = k
		for start_item.Next() != nil && start > 0 {
			int_r := start_item.Next().Value.(int)
			ans[int_r] = first
			deleted[int_r] = true

			ll.Remove(start_item.Next())
			start--
		}
		ans[start_item.Value.(int)] = first
		ll.Remove(start_item)
		if first == "1" {
			first = "2"
		} else {
			first = "1"
		}
	}
	for i := 0; i < n; i++ {
		write(f, ans[a[i]])
	}
	write(f, "\n")

}
