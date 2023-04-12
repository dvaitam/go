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
func other(i int, j int) int {
	m := map[int]bool{}
	m[i] = true
	m[j] = true
	for k := 0; k < 3; k++ {
		if !m[k] {
			return k
		}
	}
	return 0
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	c := make([][]int64, 3)
	for i := 0; i < 3; i++ {
		c[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &c[i][j])
		}
	}
	adj := make([][]int, n+1)
	ok := true
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		if len(adj[u]) > 2 || len(adj[v]) > 2 {
			ok = false
		}
	}
	if ok {
		line := make([]int, 0)
		for i := 1; i <= n; i++ {
			if len(adj[i]) == 1 {
				line = append(line, i)
				line = append(line, adj[i][0])
				break
			}
		}

		for len(adj[line[len(line)-1]]) > 1 {
			last := line[len(line)-1]
			for _, v := range adj[last] {
				if v == line[len(line)-2] {
					continue
				} else {
					line = append(line, v)
					break
				}
			}
		}
		adj = nil
		min_cost := int64(-1)
		ans_i, ans_j := -1, -1
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == j {
					continue
				}
				cost := int64(0)
				curr, nxt := i, j
				for _, k := range line {
					cost += c[curr][k-1]
					curr, nxt = nxt, other(curr, nxt)
				}
				if min_cost == -1 || cost < min_cost {
					min_cost = cost
					ans_i, ans_j = i, j
				}
			}
		}
		c = nil
		curr, nxt := ans_i, ans_j
		write(f, min_cost, "\n")
		ans := make([]int, n+1)
		for _, k := range line {
			ans[k] = curr + 1
			curr, nxt = nxt, other(curr, nxt)
		}
		line = nil
		for i := 1; i <= n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
	} else {
		write(f, "-1\n")
	}
}
