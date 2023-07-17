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
		diff := a[0] - b[0]
		for i := 0; i < n; i++ {
			a[i] -= diff
		}
		write(f, a, b, "\n")
		ok := true
		for i := 1; i < n; i++ {
			if 2*b[i]+h > 2*a[i]+w {
				if b[i]+h > a[i]+w {
					write(f, "1 failing", i, "\n")
					ok = false
					break
				}
			} else if 2*b[i]-h < 2*a[i]-w {
				if a[i]-w > b[i]-h {
					write(f, "2  failing", i, "\n")
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
