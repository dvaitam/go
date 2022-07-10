// 1702E. Split Into Two Sets
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dsu struct {
	p []int
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
		d.p[x] = y
		return true
	}
	return false
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, a, b int
		fmt.Fscan(reader, &n)
		possible := true
		deg := make([]int, n)
		dsu := Dsu{}
		dsu.p = make([]int, 2*n)
		for i := 0; i < 2*n; i++ {
			dsu.p[i] = i
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a, &b)
			a--
			b--
			deg[a]++
			deg[b]++
			dsu.unite(a, b+n)
			dsu.unite(a+n, b)
		}
		for i := 0; i < n; i++ {
			if deg[i] > 2 {
				possible = false
			}
			if dsu.get(i) == dsu.get(i+n) {
				possible = false
			}
		}
		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
