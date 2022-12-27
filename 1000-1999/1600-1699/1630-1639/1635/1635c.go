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
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		ok := true
		for i := 1; i < n; i++ {
			if a[i] < a[i-1] {
				ok = false
				break
			}
		}
		if ok {
			write(f, "0\n")
		} else {
			if (a[n-2] < 0 && a[n-1] < 0) || a[n-2] > a[n-1] {
				write(f, "-1\n")
			} else {
				write(f, n-2, "\n")
				for i := 1; i <= n-2; i++ {
					write(f, i, n-1, n, "\n")
				}
			}
		}

	}
}
