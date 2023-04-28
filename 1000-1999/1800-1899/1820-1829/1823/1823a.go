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
		total := (n * (n - 1)) / 2
		possible := false
		for i := 0; i <= n/2; i++ {
			poss := total - i*(n-i)
			if poss == k {
				possible = true
				write(f, "YES\n")
				for j := 0; j < i; j++ {
					write(f, "-1 ")
				}
				for j := i; j < n; j++ {
					write(f, "1 ")
				}
				write(f, "\n")
				break
			}
		}
		if !possible {
			write(f, "NO\n")
		}
	}
}
