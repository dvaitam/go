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
func get_beauty(root int, adj map[int][]int, parent int, l []int64, r []int64) []int64 {
	ans := make([]int64, 2)
	for i := 0; i < len(adj[root]); i++ {
		ch := adj[root][i]
		if ch == parent {
			continue
		}
		chb := get_beauty(ch, adj, root, l, r)
		ans[0] += max(chb[0]+abs(l[ch]-l[root]), chb[1]+abs(r[ch]-l[root]))
		ans[1] += max(chb[0]+abs(l[ch]-r[root]), chb[1]+abs(r[ch]-r[root]))
	}
	return ans
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
		l := make([]int64, n)
		r := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &l[i], &r[i])
		}
		adj := map[int][]int{}
		for i := 0; i < n-1; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			u--
			v--
			if adj[v] == nil {
				adj[v] = make([]int, 0)
			}

			if adj[u] == nil {
				adj[u] = make([]int, 0)
			}
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
		//write(f, adj, "\n")
		ans := get_beauty(0, adj, 0, l, r)
		write(f, max(ans[0], ans[1]), "\n")
	}
}
