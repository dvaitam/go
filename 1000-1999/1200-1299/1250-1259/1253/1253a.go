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
		l := 0
		r := n - 1
		for l < n && a[l] == b[l] {
			l++
		}
		for r >= 0 && a[r] == b[r] {
			r--
		}
		ok := true
		if l <= r {
			diff := b[l] - a[l]
			if diff < 0 {
				ok = false
			} else {
				for k := l; k <= r; k++ {
					if b[k]-a[k] != diff {
						ok = false
						break
					}
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
