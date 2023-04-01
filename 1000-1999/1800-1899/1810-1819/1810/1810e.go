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

type Pair struct {
	index, danger, vertex int
}
type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].danger < pq[j].danger
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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([]int, n+1)
		possibles := make([]int, 0)
		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] == 0 {
				possibles = append(possibles, i)
			}
		}
		adj := map[int][]int{}
		for i := 0; i < m; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			if adj[u] == nil {
				adj[u] = make([]int, 0)
			}
			if adj[v] == nil {
				adj[v] = make([]int, 0)
			}
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
		ok := false
		tried := map[int]bool{}
		//write(f, possibles, "\n")
		for i := 0; i < len(possibles); i++ {
			if tried[possibles[i]] {
				continue
			}
			pq := make(PriorityQueue, 0)
			health := 0
			heap.Push(&pq, &Pair{danger: 0, vertex: possibles[i]})
			included := map[int]bool{}
			included[possibles[i]] = true
			for pq.Len() > 0 {
				if health < pq[0].danger {
					break
				} else {
					item := heap.Pop(&pq).(*Pair)
					if item.danger == 0 {
						tried[item.vertex] = true
					}
					for _, v := range adj[item.vertex] {
						if !included[v] {
							included[v] = true
							heap.Push(&pq, &Pair{danger: a[v], vertex: v})
						}
					}

					health++
				}
			}
			if pq.Len() == 0 {
				if len(included) == n {
					ok = true
				}
				break
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
