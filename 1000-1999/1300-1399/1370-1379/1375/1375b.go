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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([][]int, n)
		ok := true
		for i := 0; i < n; i++ {
			a[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(reader, &a[i][j])
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				count := 0
				if i-1 >= 0 {
					count++
				}
				if j-1 >= 0 {
					count++
				}
				if i+1 < n {
					count++
				}
				if j+1 < m {
					count++
				}
				if a[i][j] > count {
					ok = false
				} else {
					a[i][j] = count
				}
			}
		}
		if ok {
			write(f, "YES\n")
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					write(f, a[i][j], " ")
				}
				write(f, "\n")
			}
		} else {
			write(f, "NO\n")
		}

	}
}
