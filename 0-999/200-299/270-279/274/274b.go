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
func go_down(root int, parent int, adj map[int][]int, v []int64) (int64, int64) {
	curr, dec, inc := v[root], int64(0), int64(0)
	for _, val := range adj[root] {
		if val == parent {
			continue
		}
		d, i := go_down(val, root, adj, v)
		dec = max(dec, d)
		inc = max(inc, i)
	}
	//fmt.Println("before", root, curr, dec, inc)
	curr = curr - dec + inc
	//fmt.Println(root, curr, dec, inc)
	if curr > 0 {
		dec += curr
	} else {
		inc -= curr
	}
	return dec, inc
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	adj := map[int][]int{}
	first := -1
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		if first == -1 {
			first = u
		}
		if adj[u] == nil {
			adj[u] = make([]int, 0)
		}
		if adj[v] == nil {
			adj[v] = make([]int, 0)
		}
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	v := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &v[i])
	}
	//	write(f, v, "\n")
	dec, inc := go_down(1, 1, adj, v)
	// if first == 80854 {
	// 	write(f, dec, inc, "\n")
	// }
	write(f, dec+inc, "\n")
}
