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

type Item struct {
	index    int
	priority int64
	vertex   int64
	x, y     int64
	on_queue bool
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
func (pq *PriorityQueue) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Point struct {
	x, y int64
}
type Segment struct {
	a, b   int64
	points []int64
}

var mod int64

func add_edges(rsegments map[int64][]Segment,
	row int64, dest_row int64,
	adj map[int64]map[int64]int64,
	adjm map[int64][]int64,
	col int64) {
	pl := len(rsegments[dest_row])
	for sp := col - 1; sp <= col+1; sp++ {
		j := sort.Search(pl, func(i int) bool {
			return rsegments[dest_row][i].b >= sp
		})
		if j < pl && sp >= rsegments[dest_row][j].a {
			source, dest := row*mod+col, (dest_row)*mod+sp
			if adj[source] == nil {
				adj[source] = map[int64]int64{}
			}
			if adj[dest] == nil {
				adj[dest] = map[int64]int64{}
			}
			if adj[source][dest] == 0 {
				rsegments[dest_row][j].points = append(rsegments[dest_row][j].points, sp)
				adj[source][dest] = 1
				adjm[source] = append(adjm[source], dest)
			}
			if adj[dest][source] == 0 {
				adj[dest][source] = 1
				adjm[dest] = append(adjm[dest], source)

			}
		}
	}

}

func main() {
	mod = int64(1000000001)
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var x0, y0, x1, y1 int64
	fmt.Fscan(reader, &x0, &y0, &x1, &y1)
	var n int
	fmt.Fscan(reader, &n)
	m := map[int64]bool{}
	for i := 0; i < n; i++ {
		var r, a, b int64
		fmt.Fscan(reader, &r, &a, &b)
		for y := a; y <= b; y++ {
			m[r*mod+y] = true
		}
	}
	dist := map[int64]*Item{}
	for k := range m {
		dist[k] = &Item{vertex: k, priority: mod, x: k / mod, y: k % mod}
	}
	source := x0*mod + y0
	dist[source].priority = 0
	dist[source].on_queue = true
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, dist[source])
	for pq.Len() > 0 {
		top := heap.Pop(&pq).(*Item)
		top.on_queue = false
		for j := top.y - 1; j <= top.y+1; j++ {
			for i := top.x - 1; i <= top.x+1; i++ {
				dest := i*mod + j
				if !m[dest] {
					continue
				}
				if top.priority+1 < dist[dest].priority {
					if dist[dest].on_queue {
						pq.update(dist[dest], top.priority+1)
					} else {
						dist[dest].on_queue = true
						dist[dest].priority = top.priority + 1
						heap.Push(&pq, dist[dest])
					}
				}
			}
		}
	}
	dest := x1*mod + y1
	if dist[dest].priority != mod {
		write(f, dist[dest].priority, "\n")
	} else {
		write(f, "-1\n")
	}

}
