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
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		start := 0
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				start = i
				break
			}
		}
		end := 0
		for i := n - 1; i >= 0; i-- {
			if a[i] != b[i] {
				end = i
				break
			}
		}
		mi, mx := a[start], a[end]
		for i := start; i <= end; i++ {
			mi = min(mi, a[i])
			mx = max(mx, a[i])
		}
		for start > 0 && a[start-1] <= mi && a[start] >= a[start-1] {
			start--
		}
		for end+1 < n && a[end+1] >= mx && a[end] <= a[end+1] {
			end++
		}
		write(f, start+1, end+1, "\n")
	}
}
