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
		m := map[int]bool{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = true
		}
		mx := 0
		for m[mx] {
			mx++
		}
		if mx == n {
			write(f, "NO\n")
		} else {
			if !m[mx+1] {
				write(f, "YES\n")
			} else {
				nm := map[int]bool{}
				for i := 0; i < n; i++ {
					if a[i] != mx+1 {
						nm[a[i]] = true
					} else {
						break
					}
				}
				for i := n - 1; i >= 0; i-- {
					if a[i] != mx+1 {
						nm[a[i]] = true
					} else {
						break
					}
				}
				nmx := 0
				for nm[nmx] {
					nmx++
				}
				if nmx == mx {
					write(f, "Yes\n")
				} else {
					write(f, "No\n")
				}
			}
		}

	}
}
