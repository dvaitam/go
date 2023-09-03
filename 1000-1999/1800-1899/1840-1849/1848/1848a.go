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
		var n, m, k int
		fmt.Fscan(reader, &n, &m, &k)
		x := make([]int, k+1)
		y := make([]int, k+1)
		for i := 0; i < k+1; i++ {
			fmt.Fscan(reader, &x[i], &y[i])
		}
		if k < min(min(n, m), 4) {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
