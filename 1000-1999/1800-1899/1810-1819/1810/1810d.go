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
func get_days(h int64, a int64, b int64) int64 {
	if h <= a {
		return 1
	} else {
		rem := h - a
		if rem%(a-b) == 0 {
			return (rem / (a - b)) + 1
		} else {
			return (rem / (a - b)) + 2
		}
	}
}
func get_h(n int64, a int64, b int64) int64 {
	return (n-1)*(a-b) + a
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var q int
		fmt.Fscan(reader, &q)
		mih, mxh := int64(0), int64(1000000000*1000000000)
		for i := 0; i < q; i++ {
			var tp, a, b, n int64
			fmt.Fscan(reader, &tp)
			if tp == 1 {
				fmt.Fscan(reader, &a, &b, &n)
				pi, px := get_h(n-1, a, b)+1, get_h(n, a, b)
				if n == 1 {
					pi = 1
				}
				// write(f, "heights ", pi, px, "\n ")
				ppi, ppx := max(mih, pi), min(mxh, px)
				if ppx < ppi {
					write(f, "0 ")
				} else {
					write(f, "1 ")
					mih, mxh = ppi, ppx
				}

			} else {
				fmt.Fscan(reader, &a, &b)
				d1 := get_days(mih, a, b)
				d2 := get_days(mxh, a, b)
				if d1 != d2 {
					write(f, "-1 ")
				} else {
					write(f, d1, " ")
				}
			}
		}
		write(f, "\n")
	}
}
