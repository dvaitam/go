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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	h := make([]bool, n+1)
	v := make([]bool, n+1)
	for i := 0; i < n*n; i++ {
		var h1, v1 int
		fmt.Fscan(reader, &h1, &v1)
		if !h[h1] && !v[v1] {
			write(f, i+1, " ")
			h[h1] = true
			v[v1] = true
		}
	}
	write(f, "\n")
}
