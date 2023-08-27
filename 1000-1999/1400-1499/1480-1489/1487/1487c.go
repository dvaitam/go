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
		var n int
		fmt.Fscan(reader, &n)
		points := (3 * n * (n - 1)) / 2
		each := points / n
		p := make([]int, n+1)
		for i := 1; i < n; i++ {
			for j := i + 1; j <= n; j++ {
				if p[i]+3 <= each {
					p[i] += 3
					write(f, "1 ")
				} else if p[i]+1 <= each {
					p[i] += 1
					write(f, "0 ")
				} else {
					p[j] += 3
					write(f, "-1 ")
				}
			}
		}
		write(f, "\n")
	}
}
