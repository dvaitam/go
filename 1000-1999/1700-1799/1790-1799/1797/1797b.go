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
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Fscan(reader, &a[i][j])

			}
		}
		d := 0
		r := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][j] == 1 {
					r++
				}
				if a[i][j] != a[n-1-i][n-1-j] {
					d++
				}
			}
		}
		d = d / 2
		if k >= d && n%2 == 1 || k >= d && (k-d)%2 == 0 {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
