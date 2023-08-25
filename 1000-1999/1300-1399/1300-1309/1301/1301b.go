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
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		mi, mx := 1000000000, 0
		for i := 0; i < n; i++ {
			if a[i] == -1 {
				if i > 0 && a[i-1] != -1 {
					mi = min(mi, a[i-1])
					mx = max(mx, a[i-1])
				}
				if i+1 < n && a[i+1] != -1 {
					mi = min(mi, a[i+1])
					mx = max(mx, a[i+1])
				}
			}
		}
		k := (mi + mx) / 2
		for i := 0; i < n; i++ {
			if a[i] == -1 {
				a[i] = k
			}
		}
		m := 0
		for i := 1; i < n; i++ {
			m = max(m, abs(a[i]-a[i-1]))
		}
		write(f, m, k, "\n")

	}
}
