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
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	c := make([]int64, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			c[i] = int64(a[i])
		} else {
			c[i] = c[i-1] + int64(a[i])
		}
	}
	ans := make([]int64, n)
	for i := 0; i < m; i++ {
		ans[i] = c[i]
	}
	for i := m; i < n; i++ {
		ans[i] = ans[i-m] + c[i]
	}
	for i := 0; i < n; i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
