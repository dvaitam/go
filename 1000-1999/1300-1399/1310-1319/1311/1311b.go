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
type Range struct {
	i, j int
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
func equal(a []int, b []int) bool {
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([]int, n)
		sa := make([]int, n)
		b := make([]bool, n)
		p := make([]int, m)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			sa[i] = a[i]
		}
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &p[i])
		}
		sort.Ints(p)
		for i := 0; i < m; i++ {
			if p[i]-1 < n {
				b[p[i]-1] = true
			}
			if p[i] < n {
				b[p[i]] = true
			}
		}
		start := p[0]
		segments := make([]Range, 0)
		for i := 1; i < m; i++ {
			if p[i] != p[i-1]+1 {
				segments = append(segments, Range{i: start - 1, j: p[i-1]})
				start = p[i]
			}
		}
		segments = append(segments, Range{i: start - 1, j: p[m-1]})
		sort.Ints(sa)
		ok := true
		for i := 0; i < n; i++ {
			if !b[i] && a[i] != sa[i] {
				ok = false
				break
			}
		}
		if ok {
			for i := 0; i < len(segments); i++ {
				start, end := segments[i].i, segments[i].j+1
				if !equal(a[start:end], sa[start:end]) {
					ok = false
					break
				}
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
