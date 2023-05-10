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

type Edge struct {
	u, v, t int
}

type Item struct {
	index    int
	priority int
	vertex   int
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
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	var n, m, T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &T)
	edges := make([]Edge, 0)
	//adj := make([][]int, n+1)
	edu := map[int]int{}
	for i := 0; i < m; i++ {
		var u, v, t int
		fmt.Fscan(reader, &u, &v, &t)
		edges = append(edges, Edge{u, v, t})
		//	adj[v] = append(adj[v], u)
		edu[v*10000+u] = t
	}
	// dist := map[int]*Item{}
	// for i := 1; i <= n; i++ {
	// 	dist[i] = &Item{vertex: i, priority: T}
	// }
	// source := n
	// dist[source].priority = 0
	// dist[source].on_queue = true
	// pq := make(PriorityQueue, 0)
	// heap.Init(&pq)
	// heap.Push(&pq, dist[source])
	// for pq.Len() > 0 {
	// 	top := heap.Pop(&pq).(*Item)
	// 	top.on_queue = false
	// 	for _, dest := range adj[top.vertex] {
	// 		if top.priority+edu[top.vertex*10000+dest] < dist[dest].priority {
	// 			if dist[dest].on_queue {
	// 				pq.update(dist[dest], top.priority+edu[top.vertex*10000+dest])
	// 			} else {
	// 				dist[dest].on_queue = true
	// 				dist[dest].priority = top.priority + edu[top.vertex*10000+dest]
	// 				heap.Push(&pq, dist[dest])
	// 			}
	// 		}
	// 	}
	// }

	parent := make([]int, n+1)
	dp := make([]int, n+1)
	nodes := map[int]map[int]int{} // make([][]int, n+1)
	for i := 1; i <= n; i++ {
		nodes[i] = map[int]int{}
	}

	nodes[1][1] = 0
	for i := 0; i < m; i++ {
		if edges[i].u == 1 {
			dp[edges[i].v] = edges[i].t
			parent[edges[i].v] = 1
			nodes[2][edges[i].v] = edges[i].t
			//paths[edges[i].v] = append(paths[edges[i].v], 1)
		}
	}
	//write(f, parent, "\n")
	ans := 2
	//path := make([]int, 0)
	for j := 3; j <= n; j++ {
		nxt_dp := make([]int, n+1)
		for i := 0; i < m; i++ {
			if dp[edges[i].u] != 0 && dp[edges[i].u]+edges[i].t <= T {

				//	if dp[edges[i].u] != 0 && dp[edges[i].u]+edges[i].t+dist[edges[i].v].priority <= T {
				if nxt_dp[edges[i].v] == 0 || nxt_dp[edges[i].v] > dp[edges[i].u]+edges[i].t {
					nxt_dp[edges[i].v] = dp[edges[i].u] + edges[i].t
					parent[edges[i].v] = edges[i].u
					nodes[j][edges[i].v] = dp[edges[i].u] + edges[i].t
					if edges[i].v == n {
						ans = max(ans, j)
					}
				}

			}
		}
		// for i := 1; i <= n; i++ {
		// 	if nxt_dp[i] != 0 {
		// 		curr := paths[parent[i]]
		// 		//copy(curr, paths[parent[i]])
		// 		curr = append(curr, parent[i])
		// 		paths[i] = curr
		// 	}
		// }
		// for i := 0; i <= n; i++ {
		// 	if nxt_dp[i] != 0 && parent[i] != n {
		// 		paths[parent[i]] = nil
		// 	}
		// }
		dp = nxt_dp
		//write(f, paths, "\n")

	}
	//write(f, nodes, "\n")
	//write(f, parent, "\n")
	write(f, ans, "\n")
	route := make([]int, 0)
	route = append(route, n)
	for i := ans; i >= 1; i-- {
		curr := route[len(route)-1]
		nxt_node := -1
		for j := 1; j <= n; j++ {
			if edu[curr*10000+j] > 0 {
				if nodes[i-1][j] > 0 && edu[curr*10000+j]+nodes[i-1][j] == nodes[i][curr] {
					nxt_node = j
				}
			}
		}
		if nxt_node == -1 {
			nxt_node = 1
		}
		route = append(route, nxt_node)
	}
	// tmp := n
	// for tmp != 1 {
	// 	route = append(route, tmp)
	// 	tmp = parent[tmp]
	// }
	//route := paths[n]
	//route = append(route, n)
	//write(f, paths[n], "\n")
	//write(f, route, "\n")
	//route = append(route, 1)
	for i := ans - 1; i >= 0; i-- {
		write(f, route[i], " ")
	}
	write(f, "\n")

}
