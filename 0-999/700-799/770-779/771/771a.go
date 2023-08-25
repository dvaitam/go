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

type Dsu struct {
	p          []int
	sz         []int
	components int
}

func (d Dsu) get(x int) int {
	if x == d.p[x] {
		return x
	} else {
		d.p[x] = d.get(d.p[x])
		return d.p[x]
	}
}
func (d Dsu) unite(x int, y int) bool {
	x = d.get(x)
	y = d.get(y)
	if x != y {
		if d.sz[x] < d.sz[y] {
			x, y = y, x
		}
		d.sz[x] += d.sz[y]
		d.p[y] = x
		d.components--
		return true
	}
	return false
}
func (d Dsu) size(x int) int {
	return d.sz[d.get(x)]
}
func give_dsu(n int) Dsu {
	p := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		sz[i] = 1
	}

	return Dsu{p, sz, n}
}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	dsu := give_dsu(n + 1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		dsu.unite(u, v)
	}
	sizes := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sizes[dsu.get(i)]++
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if sizes[i] > 0 {
			ans += (sizes[i] * (sizes[i] - 1)) / 2
		}
	}
	if ans == m {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
