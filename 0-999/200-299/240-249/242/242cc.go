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
	segments := map[int64][]Segment{}
	for i := 0; i < n; i++ {
		var r, a, b int64
		fmt.Fscan(reader, &r, &a, &b)
		segments[r] = append(segments[r], Segment{a: a, b: b})
	}
	// segments[x0] = append(segments[x0], Segment{a: y0, b: y0})
	// segments[x1] = append(segments[x1], Segment{a: y1, b: y1})

	for k := range segments {
		sort.Slice(segments[k], func(i, j int) bool { return segments[k][i].a < segments[k][j].a })
	}
	rsegments := map[int64][]Segment{}
	for k := range segments {
		new_segments := make([]Segment, 0)

		for i := 0; i < len(segments[k]); i++ {
			nl := len(new_segments)

			if nl == 0 || segments[k][i].a > new_segments[nl-1].b {
				new_segments = append(new_segments, segments[k][i])
			} else {
				new_segments[nl-1].b = max(segments[k][i].b, new_segments[nl-1].b)
			}
		}
		rsegments[k] = new_segments
	}

	adj := map[int64]map[int64]int64{}
	adjm := map[int64][]int64{}
	for k, s := range rsegments {
		for ri := 0; ri < len(s); ri++ {
			s[ri].points = append(s[ri].points, s[ri].a)
			add_edges(rsegments, k, k-1, adj, adjm, s[ri].a)
			add_edges(rsegments, k, k+1, adj, adjm, s[ri].a)
			s[ri].points = append(s[ri].points, s[ri].b)
			add_edges(rsegments, k, k-1, adj, adjm, s[ri].b)
			add_edges(rsegments, k, k+1, adj, adjm, s[ri].b)
		}
	}
	add_edges(rsegments, x0, x0-1, adj, adjm, y0)
	add_edges(rsegments, x0, x0+1, adj, adjm, y0)
	add_edges(rsegments, x1, x1-1, adj, adjm, y1)
	add_edges(rsegments, x1, x1+1, adj, adjm, y1)
	for x := x0; x >= 0; x-- {
		pl := len(rsegments[x])
		j := sort.Search(pl, func(i int) bool {
			return rsegments[x][i].b >= y0
		})
		if j < pl && y0 >= rsegments[x][j].a {
			rsegments[x][j].points = append(rsegments[x][j].points, y0)
			if x < x0 {
				source, dest := x*mod+y0, (x+1)*mod+y0
				if adj[source] == nil {
					adj[source] = map[int64]int64{}
				}
				if adj[dest] == nil {
					adj[dest] = map[int64]int64{}
				}
				if adj[source][dest] == 0 {
					adj[source][dest] = 1
					adjm[source] = append(adjm[source], dest)
				}
				if adj[dest][source] == 0 {
					adj[dest][source] = 1
					adjm[dest] = append(adjm[dest], source)
				}

			}
		} else {
			break
		}
	}
	for x := x0; true; x++ {
		pl := len(rsegments[x])
		j := sort.Search(pl, func(i int) bool {
			return rsegments[x][i].b >= y0
		})
		if j < pl && y0 >= rsegments[x][j].a {
			rsegments[x][j].points = append(rsegments[x][j].points, y0)
			if x > x0 {
				source, dest := x*mod+y0, (x-1)*mod+y0
				if adj[source] == nil {
					adj[source] = map[int64]int64{}
				}
				if adj[dest] == nil {
					adj[dest] = map[int64]int64{}
				}
				if adj[source][dest] == 0 {
					adj[source][dest] = 1
					adjm[source] = append(adjm[source], dest)
				}
				if adj[dest][source] == 0 {
					adj[dest][source] = 1
					adjm[dest] = append(adjm[dest], source)
				}

			}
		} else {
			break
		}
	}

	for x := x1; x >= 0; x-- {
		pl := len(rsegments[x])
		j := sort.Search(pl, func(i int) bool {
			return rsegments[x][i].b >= y1
		})
		if j < pl && y1 >= rsegments[x][j].a {
			rsegments[x][j].points = append(rsegments[x][j].points, y1)
			if x < x1 {
				source, dest := x*mod+y1, (x+1)*mod+y1
				if adj[source] == nil {
					adj[source] = map[int64]int64{}
				}
				if adj[dest] == nil {
					adj[dest] = map[int64]int64{}
				}
				if adj[source][dest] == 0 {
					adj[source][dest] = 1
					adjm[source] = append(adjm[source], dest)
				}
				if adj[dest][source] == 0 {
					adj[dest][source] = 1
					adjm[dest] = append(adjm[dest], source)
				}

			}
		} else {
			break
		}
	}
	for x := x1; true; x++ {
		pl := len(rsegments[x])
		j := sort.Search(pl, func(i int) bool {
			return rsegments[x][i].b >= y1
		})
		if j < pl && y1 >= rsegments[x][j].a {
			rsegments[x][j].points = append(rsegments[x][j].points, y1)
			if x > x1 {
				source, dest := x*mod+y1, (x-1)*mod+y1
				if adj[source] == nil {
					adj[source] = map[int64]int64{}
				}
				if adj[dest] == nil {
					adj[dest] = map[int64]int64{}
				}
				if adj[source][dest] == 0 {
					adj[source][dest] = 1
					adjm[source] = append(adjm[source], dest)
				}
				if adj[dest][source] == 0 {
					adj[dest][source] = 1
					adjm[dest] = append(adjm[dest], source)
				}

			}
		} else {
			break
		}
	}

	pl := len(rsegments[x1])
	j := sort.Search(pl, func(i int) bool {
		return rsegments[x1][i].b >= y1
	})
	if y0 >= rsegments[x1][j].a {
		rsegments[x1][j].points = append(rsegments[x1][j].points, y1)
	}
	for k, s := range rsegments {
		for ri := 0; ri < len(s); ri++ {
			sort.Slice(s[ri].points, func(i, j int) bool { return s[ri].points[i] < s[ri].points[j] })
			points := s[ri].points
			for i := 1; i < len(points); i++ {
				if points[i] > points[i-1] {
					source, dest := k*mod+points[i-1], (k)*mod+points[i]
					if adj[source] == nil {
						adj[source] = map[int64]int64{}
					}
					if adj[dest] == nil {
						adj[dest] = map[int64]int64{}
					}
					if adj[source][dest] == 0 {
						adj[source][dest] = points[i] - points[i-1]
						adjm[source] = append(adjm[source], dest)
					}
					if adj[dest][source] == 0 {
						adj[dest][source] = points[i] - points[i-1]
						adjm[dest] = append(adjm[dest], source)
					}
				}
			}
		}
	}
	dist := map[int64]*Item{}
	for k := range adj {
		dist[k] = &Item{vertex: k, priority: mod}
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
		for v, d := range adj[top.vertex] {
			if top.priority+d < dist[v].priority {
				if dist[v].on_queue {
					pq.update(dist[v], top.priority+d)
				} else {
					dist[v].on_queue = true
					dist[v].priority = top.priority + d
					heap.Push(&pq, dist[v])
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
