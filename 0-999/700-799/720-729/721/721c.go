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
var mm map[int64]int
var mmt map[int64]int64

func backtrack(curr *list.List, adj map[int64][]Edge, t int64, T int64, n int64) {
	if t > T {
		return
	}
	if int64(len(ans)) == n {
		return
	}
	last := curr.Back().Value.(int64)
	if curr.Len() <= mm[last] && t > mmt[last] {
		return
	} else {
		if curr.Len() >= mm[last] && t <= mmt[last] || mmt[last] == 0 {
			mm[last] = curr.Len()
			mmt[last] = t
		}
	}

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
	mm = map[int64]int{}
	mmt = map[int64]int64{}
	for i := int64(0); i < m; i++ {
		var u, v, t int64
		fmt.Fscan(reader, &u, &v, &t)
		if adj[u] == nil {
			adj[u] = make([]Edge, 0)
		}
		adj[u] = append(adj[u], Edge{v: v, t: t})
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
