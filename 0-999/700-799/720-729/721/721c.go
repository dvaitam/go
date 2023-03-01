// 00
package main

import (
	"bufio"
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
type Edge struct {
	v, t int64
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
func reverse(ee []Edge) {
	for i, j := 0, len(ee)-1; i < j; i, j = i+1, j-1 {
		ee[i], ee[j] = ee[j], ee[i]
	}
}

var ans []int64

func backtrack(curr *list.List, adj map[int64][]Edge, t int64, T int64, n int64) {
	if t > T {
		return
	}
	if int64(len(ans)) == n {
		return
	}
	last := curr.Back().Value.(int64)

	if last == n {
		if curr.Len() > len(ans) {
			ans = make([]int64, 0)
			for e := curr.Front(); e != nil; e = e.Next() {
				ans = append(ans, e.Value.(int64))
			}
		}
		return
	}

	for _, edge := range adj[last] {
		curr.PushBack(edge.v)
		backtrack(curr, adj, t+edge.t, T, n)
		curr.Remove(curr.Back())
	}
}
func main() {
	var n, m, T int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &T)
	adj := map[int64][]Edge{}
	adjm := map[int64]map[int64]int64{}
	adjr := map[int64]map[int64]bool{}
	out_degree := map[int64]int64{}
	for i := int64(0); i < m; i++ {
		var u, v, t int64
		fmt.Fscan(reader, &u, &v, &t)
		// if adj[u] == nil {
		// 	adj[u] = make([]Edge, 0)
		// }
		if adjr[v] == nil {
			adjr[v] = map[int64]bool{}
		}
		if adjm[u] == nil {
			adjm[u] = map[int64]int64{}
		}
		adjr[v][u] = true
		//	adj[u] = append(adj[u], Edge{v: v, t: t})
		adjm[u][v] = t
		out_degree[u]++
	}
	start := make([]int64, 0)
	for i := int64(1); i <= n; i++ {
		if out_degree[i] == 0 && adjr[i] != nil && i != 1 && i != n {
			start = append(start, i)
		}
	}
	for len(start) > 0 {
		top := start[len(start)-1]
		start = start[:len(start)-1]
		if top == n {
			continue
		}
		for k := range adjr[top] {
			delete(adjm[k], top)
			out_degree[k]--
			if out_degree[k] == 0 {
				start = append(start, k)
			}
		}
	}
	for k := range adjm {
		adj[k] = make([]Edge, 0)
		for k1, t1 := range adjm[k] {
			adj[k] = append(adj[k], Edge{v: k1, t: t1})
		}
	}

	// for k := range adj {
	// 	reverse(adj[k])
	// }
	// if T == 126917982 {
	// 	T = T / 1000
	// 	write(f, len(adj[1]), "\n")
	// }
	ll := list.New()
	ll.PushBack(int64(1))

	backtrack(ll, adj, 0, T, n)
	write(f, len(ans), "\n")
	for i := 0; i < len(ans); i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")

}
