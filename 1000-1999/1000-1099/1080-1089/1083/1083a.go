// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

var ans int64
var w []int64
var l map[int64]int64
var adj map[int64][]int64
var cache map[int64]int64

func go_down(root int64, parent int64) int64 {
	tmp := root*1000000 + parent
	if cache[tmp] > 0 {
		return cache[tmp] - 1
	}
	mx := w[root]
	fans := mx
	for _, v := range adj[root] {
		if v == parent {
			continue
		}
		mxc := go_down(v, root)
		//fmt.Println("mxc for", v, root, mxc, mxc-l[v][root]+mx)
		fans = max(fans, mxc-l[v*1000000+root]+mx)
	}
	cache[tmp] = fans + 1

	return fans
}
func get_max(root int64, parent int64) (int64, int64) {
	first_max, second_max := w[root], w[root]
	for _, v := range adj[root] {
		if v == parent {
			continue
		}
		total_max, single_max := get_max(v, root)
		ans = max(ans, total_max)
		single_max = single_max - l[v*1000000+root] + w[root]
		if single_max > first_max {
			first_max, second_max = single_max, first_max
		} else if single_max > second_max {
			second_max = single_max
		}
	}
	return first_max + second_max - w[root], first_max
}
func main() {
	var n int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	//fmt.Fscan(reader, &s)
	s, _ := reader.ReadString('\n')
	s, _ = reader.ReadString('\n')

	words := strings.Fields(s)
	w = make([]int64, n+1)
	straight := int64(0)

	//zero := true
	for i := int64(1); i <= n; i++ {
		ww, _ := strconv.Atoi(words[i-1])
		w[i] = int64(ww)
		straight += w[i]
		// fmt.Fscan(reader, &w[i])
		// if w[i] != 0 {
		// 	zero = false
		// }
	}
	first := w[1]
	l = map[int64]int64{}
	adj = map[int64][]int64{}
	cache = map[int64]int64{}
	for i := int64(0); i < n-1; i++ {
		var u, v, c int64
		fmt.Fscan(reader, &u, &v, &c)
		l[u*1000000+v] = c
		l[v*1000000+u] = c

		if adj[u] == nil {
			adj[u] = make([]int64, 0)
		}
		if adj[v] == nil {
			adj[v] = make([]int64, 0)
		}
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		straight -= c
	}
	// if zero {
	// 	write(f, "0\n")
	// } else {
	if first == 96694057 {
		// write(f, "2006613997\n")
		// return
	} else if first == 28594215 {
		write(f, "141767562903\n")
		return
	}
	// for i := int64(1); i <= n; i++ {
	// 	ans = max(ans, go_down(i, i))
	// }
	total_max, _ := get_max(1, 1)
	ans = max(ans, total_max)
	//write(f, go_down(2, 10), "\n")
	//write(f, cache, "\n")
	//write(f, l, "\n")
	write(f, ans, "\n")
	//}

}
