// 00
package main

import (
	"bufio"
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}

type Edge struct {
	v int
	w int64
}

func longest(x int, parent int, adj map[int][]Edge) int64 {
	lng := int64(0)
	for _, k := range adj[x] {
		if k.v == parent {
			continue
		}
		lng = max(lng, k.w+longest(k.v, x, adj))
	}
	return lng
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	adj := map[int][]Edge{}
	ans := int64(0)
	for i := 0; i < n-1; i++ {
		var u, v int
		var w int64
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
		ans += 2 * w
	}
	ans -= longest(1, 0, adj)
	write(f, ans, "\n")
}
