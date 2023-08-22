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

var sm float64

func go_down(root int, parent int, l int, p float64, adj [][]int) {
	//fmt.Println("ye", l, p)
	if len(adj[root]) == 1 {
		sm += float64(l) * p
	}
	for _, v := range adj[root] {
		if v == parent {
			continue
		}
		prob := p / float64(len(adj[root])-1)
		if root == 1 {
			prob = p / float64(len(adj[root]))
		}
		go_down(v, root, l+1, prob, adj)
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	adj := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	go_down(1, 0, 0, 1, adj)

	write(f, sm, "\n")
}
