// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Edge struct {
	w, u, v int
}
type Pair struct {
	q, i int
}

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	e := make([]Edge, n-1)
	for i := 0; i < n-1; i++ {
		var w, u, v int
		fmt.Fscan(reader, &u, &v, &w)
		u--
		v--
		e[i] = Edge{w, u, v}
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].w < e[j].w || e[i].w == e[j].w && e[i].u < e[j].u || e[i].w == e[j].w && e[i].u == e[j].u && e[i].v < e[j].v
	})
	qs := make([]int, m)
	ord := make([]Pair, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &qs[i])
		ord[i] = Pair{qs[i], i}
	}
	sort.Slice(ord, func(i, j int) bool { return ord[i].q < ord[j].q })
	//write(f, ord, "\n")
	ans := make([]int64, m)
	dsu := give_dsu(n)
	curr := int64(0)
	i := 0
	for _, j := range ord {
		for i < n-1 && e[i].w <= qs[j.i] {
			curr += int64(dsu.size(e[i].u)) * int64(dsu.size(e[i].v))
			dsu.unite(e[i].u, e[i].v)
			i++
		}
		ans[j.i] = curr
	}
	//write(f, dsu, "\n")
	for i := 0; i < m; i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
