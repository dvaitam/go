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
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		m := map[int]bool{}
		all := make([]int, 0)
		next_index := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if !m[a[i]] {
				all = append(all, i+1)
			}
			next_index[i+1] = len(all)
			m[a[i]] = true
		}
		if len(m) == 1 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
			dsu := Dsu{}
			dsu.p = make([]int, n+1)
			for i := 0; i < n+1; i++ {
				dsu.p[i] = i
			}
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					ed1, ed2 := i+1, j+1
					if a[i] != a[j] && dsu.unite(ed1, ed2) {
						write(f, ed1, " ", ed2, "\n")
					}
				}

			}
		}

	}
}
