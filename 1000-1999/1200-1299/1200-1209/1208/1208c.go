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
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	curr := 0
	row := true
	for i := 0; i < n; i += 2 {
		col := row
		for j := 0; j < n; j += 2 {
			if col {
				a[i][j] = curr
				a[i][j+1] = curr + 1
				a[i+1][j] = curr + 2
				a[i+1][j+1] = curr + 3
			} else {
				a[i][j] = curr
				a[i][j+1] = curr + 2
				a[i+1][j] = curr + 1
				a[i+1][j+1] = curr + 3
			}
			curr += 4
			col = !col
		}
		row = !row
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			write(f, a[i][j], " ")
		}
		write(f, "\n")
	}
}
