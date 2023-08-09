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

var ans, m int

func go_down(root int, parent int, adj map[int][]int, a []int, cons int, maxcons int) {
	maxc := maxcons
	if a[root] == 1 {
		cons++
		maxc = max(maxc, cons)
	} else {
		maxc = max(maxc, cons)
		cons = 0
	}
	for _, edge := range adj[root] {
		if edge == parent {
			continue
		}
		go_down(edge, root, adj, a, cons, maxc)
	}
	//fmt.Println(root, parent, cons, maxc, maxcons, adj)
	if len(adj[root]) == 1 && root != 1 && maxc <= m {
		ans++
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	adj := map[int][]int{}
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}
	go_down(1, 0, adj, a, 0, 0)
	write(f, ans, "\n")
}
