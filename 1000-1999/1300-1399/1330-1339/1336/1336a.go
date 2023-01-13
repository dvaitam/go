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
type Node struct {
	root, h int
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

var global_count int
var h map[int]int
var children map[int]int
var visited map[int]bool
var m map[int][]int

func get_children(root int, m map[int][]int) int {

	count := 0
	for i := 0; i < len(m[root]); i++ {
		count += get_children(m[root][i], m)
		count += 1
	}
	children[root] = count
	return count
}
func remove_backedges(root int, mm map[int][]int, parent int, height int) {
	ch := m[root]
	if mm[root] == nil {
		mm[root] = make([]int, 0)
	}
	for i := 0; i < len(ch); i++ {
		if ch[i] != parent {
			mm[root] = append(mm[root], ch[i])
			h[ch[i]] = height + 1
			remove_backedges(ch[i], mm, root, height+1)
		}
	}
	if len(mm[root]) == 0 {
		mm[root] = nil
	}

}
func get_happiness(m map[int][]int, h map[int]int, children map[int]int, root int, convoy []bool) int64 {
	happiness := int64(0)
	for i := 0; i < len(m[root]); i++ {
		c := m[root][i]
		if convoy[c] {
			happiness += int64(children[c]+1) * int64(h[c])
		} else {
			happiness += get_happiness(m, h, children, c, convoy)
		}
	}
	return happiness
}
func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	cities := make([]int, n)
	for i := 1; i <= n; i++ {
		cities[i-1] = i
	}
	convoy := make([]bool, n+1)

	weights := make([]int, n)

	h = make(map[int]int)
	children = make(map[int]int)
	visited = make(map[int]bool)
	m = make(map[int][]int)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		if m[u] == nil {
			m[u] = make([]int, 0)
		}
		m[u] = append(m[u], v)
		if m[v] == nil {
			m[v] = make([]int, 0)
		}
		m[v] = append(m[v], u)

	}

	mm := make(map[int][]int)
	remove_backedges(1, mm, 1, 0)
	get_children(1, mm)
	// write(f, m, "\n")
	// write(f, mm, "\n")
	// write(f, h, "\n")
	//get_children(1, 0)
	h[1] = 0
	for i := 0; i < n; i++ {
		weights[i] = h[cities[i]] - children[cities[i]]
	}
	sort.Slice(cities, func(i, j int) bool { return h[cities[i]]-children[cities[i]] < h[cities[j]]-children[cities[j]] })
	start := n - 1
	for k > 0 {
		convoy[cities[start]] = true
		k--
		start--
	}
	// write(f, m, h, children, cities, weights, convoy, "\n")
	write(f, get_happiness(mm, h, children, 1, convoy), "\n")

}
