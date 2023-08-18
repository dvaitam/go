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
		a := make([]int, n)
		b := make([][]int, 4)
		for i := 0; i < 4; i++ {
			b[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		if a[0] == 1 {
			b[0][0] = 1
		}
		b[1][0] = n
		b[2][0] = n
		b[3][0] = n
		for i := 1; i < n; i++ {
			cost := 0
			if a[i] == 1 {
				cost = 1
			}
			b[0][i] = min(b[2][i-1], b[3][i-1]) + cost
			b[1][i] = b[0][i-1] + cost
			b[2][i] = min(b[0][i-1], b[1][i-1])
			b[3][i] = b[2][i-1]
		}
		ans := n
		for i := 0; i < 4; i++ {
			ans = min(ans, b[i][n-1])
		}
		write(f, ans, "\n")

	}
}
