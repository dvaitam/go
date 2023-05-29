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
	v, i int
}

var reading int

func go_down(x int, i int, parent int, read int, adj map[int][]Edge) {
	reading = max(reading, read)
	for _, edge := range adj[x] {
		if edge.v == parent {
			continue
		}
		if i < edge.i {
			go_down(edge.v, edge.i, x, read, adj)
		} else {
			go_down(edge.v, edge.i, x, read+1, adj)
		}
	}

}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		adj := map[int][]Edge{}
		for i := 0; i < n-1; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			adj[u] = append(adj[u], Edge{v, i})
			adj[v] = append(adj[v], Edge{u, i})
			if u == 1 || v == 1 {

			}
		}
		reading = 1
		go_down(1, n, 0, 0, adj)
		write(f, reading, "\n")
	}
}
