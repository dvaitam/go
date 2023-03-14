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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	nk := n - 1
	for nk > 0 && b[nk] == b[nk-1]+1 {
		nk--
	}
	possible := true

	if b[nk] == 1 || b[nk] == 0 {
		if b[n-1] == n {
			write(f, "0\n")
		} else {
			st := b[n-1] + 1
			for i := 0; i < nk; i++ {
				if st < b[i] || b[i] == 0 {
					st++
					continue
				} else {
					possible = false
					break
				}
			}
			//	write(f, possible, "\n")
			if possible {
				write(f, n-b[n-1], "\n")
			}
		}
	} else {
		possible = false
	}
	if !possible {
		start := 1

		ans := 0
		for i := 0; i < n; i++ {

			if start < b[i] || b[i] == 0 {
				start++
				continue
			} else {
				ans += start - b[i] + 1
				start = b[i]
			}
		}
		write(f, ans+n, "\n")
	}

}
