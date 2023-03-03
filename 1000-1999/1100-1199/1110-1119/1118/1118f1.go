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

type RedBlue struct {
	red, blue int
}

var red_blue_counts []RedBlue

func go_down(root int, parent int, a []int, adj map[int][]int) RedBlue {
	red, blue := 0, 0
	if a[root] == 1 {
		red++
	} else if a[root] == 2 {
		blue++
	}
	for _, v := range adj[root] {
		if v == parent {
			continue
		}
		v_ans := go_down(v, root, a, adj)
		red += v_ans.red
		blue += v_ans.blue
	}
	red_blue_counts[root] = RedBlue{red: red, blue: blue}
	return red_blue_counts[root]
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n+1)
	red_blue_counts = make([]RedBlue, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	adj := map[int][]int{}
	for i := 0; i < n-1; i++ {
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
	go_down(1, 0, a, adj)
	ans := 0
	root := red_blue_counts[1]
	for i := 2; i <= n; i++ {
		curr := red_blue_counts[i]
		if curr.blue == 0 && curr.red == root.red || curr.red == 0 && curr.blue == root.blue {
			ans++
		}
	}
	write(f, ans, "\n")
}
