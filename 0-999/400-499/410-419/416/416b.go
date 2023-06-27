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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &m, &n)
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &t[i][j])
		}
	}
	for j := 1; j < n; j++ {
		t[0][j] += t[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				t[i][j] += t[i-1][j]
			} else {
				t[i][j] += max(t[i-1][j], t[i][j-1])
			}
		}
	}
	for i := 0; i < m; i++ {
		write(f, t[i][n-1], " ")
	}
	write(f, "\n")

}
