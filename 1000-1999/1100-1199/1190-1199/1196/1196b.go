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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		r := make([]int, 0)
		sm := 0
		for i := 0; i < n; i++ {
			sm += a[i]
			if sm%2 == 1 {
				if len(r) < k-1 {
					r = append(r, i+1)
					sm = 0
				}
			}
		}
		if sm%2 == 1 && len(r) == k-1 && (len(r) > 0 && r[k-2] != n || len(r) == 0 && k == 1) {
			r = append(r, n)
			write(f, "YES\n")
			for i := 0; i < k; i++ {
				write(f, r[i])
				if i == k-1 {
					write(f, "\n")
				} else {
					write(f, " ")
				}
			}
		} else {
			write(f, "NO\n")
		}
	}
}
