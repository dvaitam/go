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
func main() {
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &k)
	h := make([]bool, n+1)
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(reader, &x)
		h[x] = true
	}
	ans := 1
	curr := 1
	hole := false
	for i := 0; i < k; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)

		if curr == u {
			if h[curr] {
				ans = curr
				hole = true
			}
			if !hole {
				ans = v
				if h[v] {
					hole = true
				}
				curr = v
			}
		} else if curr == v {
			if h[curr] {
				ans = curr
				hole = true
			}
			if !hole {
				ans = u
				if h[u] {
					hole = true
				}
				curr = u
			}
		}
	}
	write(f, ans, "\n")
}
