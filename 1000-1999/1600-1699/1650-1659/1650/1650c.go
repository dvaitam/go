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
type Point struct {
	x, w, i int
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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([]Point, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &a[i].x, &a[i].w)
			a[i].i = i + 1
		}
		sort.Slice(a, func(i, j int) bool { return a[i].w < a[j].w })
		a = a[:2*n]
		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
		ans := 0
		for i := 0; i < 2*n; i++ {
			ans += a[i].w
		}
		write(f, ans, "\n")
		for i := 0; i < n; i++ {
			write(f, a[i].i, a[2*n-i-1].i, "\n")
		}
	}
}
