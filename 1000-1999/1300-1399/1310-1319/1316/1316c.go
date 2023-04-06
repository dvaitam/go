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
	var n, m, p int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &p)
	a := make([]int, n)
	b := make([]int, m)
	ai, bi := -1, -1
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i]%p != 0 && ai == -1 {
			ai = i
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
		if b[i]%p != 0 && bi == -1 {
			bi = i
		}
	}
	write(f, ai+bi, "\n")
}
