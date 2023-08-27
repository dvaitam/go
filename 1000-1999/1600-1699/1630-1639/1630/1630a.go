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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		if n == 4 && k == 3 {
			write(f, "-1\n")
		} else {
			if k == 0 {
				for x := 0; x < n/2; x++ {
					write(f, x, n-1-x, "\n")
				}
			} else if k == n-1 {
				write(f, 0, n/2-2, "\n")
				write(f, n/2-1, n-1, "\n")
				write(f, n/2, n/2+1, "\n")
				for x := 1; x < n/2-2; x++ {
					write(f, x, n-1-x, "\n")
				}
			} else {
				write(f, 0, n-1-k, "\n")
				write(f, k, n-1, "\n")
				for x := 1; x < n/2; x++ {
					if x == k || x == n-1-k {
						continue
					}
					write(f, x, n-1-x, "\n")
				}
			}
		}
	}
}
