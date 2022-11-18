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
		to := map[int]bool{}
		from := map[int]bool{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] == 0 {
				to[i] = true
			} else {
				from[i] = true
			}
		}
		ok := false
		div := -1
		for i := 0; i < n-1; i++ {
			if to[i] && from[i+1] {
				ok = true
				div = i + 1
				break
			}
		}
		if a[0] == 1 {
			write(f, n+1, " ")
			for i := 1; i <= n; i++ {
				write(f, i)
				if i == n {
					write(f, "\n")
				} else {
					write(f, " ")
				}
			}
		} else if a[n-1] == 0 {
			for i := 1; i <= n+1; i++ {
				write(f, i)
				if i == n+1 {
					write(f, "\n")
				} else {
					write(f, " ")
				}
			}
		} else if ok {
			for i := 1; i <= n; i++ {
				write(f, i)
				if i == div {
					write(f, " ", n+1)
				}
				if i == n {
					write(f, "\n")
				} else {
					write(f, " ")
				}
			}
		} else {
			write(f, "-1\n")
		}

	}
}
