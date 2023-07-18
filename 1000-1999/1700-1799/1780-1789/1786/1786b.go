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
		var n, w, h int
		fmt.Fscan(reader, &n, &w, &h)
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		diff := (a[0] - b[0])
		for i := 0; i < n; i++ {
			a[i] -= diff
		}
		for i := 0; i < n; i++ {
			b[i] += (w - h)
		}
		max_delta := (b[0] - h) - (a[0] - w)
		delta := 0

		ok := true
		for i := 1; i < n; i++ {
			if b[i]-h-delta < a[i]-w {
				ok = false
				break
			}
			if b[i]+h-delta > a[i]+w {
				delta += b[i] + h - delta - (a[i] + w)
				if delta > max_delta {
					ok = false
					break
				}
			}
			max_delta = min(max_delta, b[i]-h-(a[i]-w))
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
