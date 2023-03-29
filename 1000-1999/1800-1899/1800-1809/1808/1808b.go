// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([][]int64, m)
		for i := 0; i < m; i++ {
			a[i] = make([]int64, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(reader, &a[j][i])
			}
		}
		for k := 0; k < m; k++ {
			sort.Slice(a[k], func(i, j int) bool { return a[k][i] < a[k][j] })
		}
		ans := int64(0)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				ans += int64(j)*a[i][j] - int64(n-j-1)*a[i][j]
			}
		}
		write(f, ans, "\n")

	}
}
