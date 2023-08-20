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

type Customer struct {
	t, l, h int64
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var m int64
		fmt.Fscan(reader, &n, &m)
		c := make([]Customer, n)
		for i := 0; i < n; i++ {
			var t, l, h int64
			fmt.Fscan(reader, &t, &l, &h)
			c[i] = Customer{t: t, l: l, h: h}
		}
		sort.Slice(c, func(i, j int) bool { return c[i].t < c[j].t })
		l, h := m, m
		t := int64(0)
		ok := true
		for i := 0; i < n; i++ {
			diff := c[i].t - t
			l, h = l-diff, h+diff
			nl, nh := max(l, c[i].l), min(h, c[i].h)
			if nh < nl {
				ok = false
				break
			}
			t = c[i].t
			l, h = nl, nh
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
