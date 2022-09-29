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
		if n%2 == 1 && k != 1 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
			odd_start := 1
			even_start := 2
			for i := 0; i < n; i++ {
				for j := 0; j < k; j++ {
					if i%2 == 0 {
						write(f, odd_start)
						odd_start += 2
					} else {
						write(f, even_start)
						even_start += 2
					}
					if j == k-1 {
						write(f, "\n")
					} else {
						write(f, " ")
					}
				}
			}
		}
	}
}
