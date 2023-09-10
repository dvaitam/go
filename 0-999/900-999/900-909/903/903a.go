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
	p := make([]bool, 101)
	for i := 0; i < 35; i++ {
		for j := 0; j < 16; j++ {
			if 3*i+7*j <= 100 {
				p[3*i+7*j] = true
			}
		}
	}
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		if p[x[i]] {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
