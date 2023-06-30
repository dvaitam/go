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
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		if a[n-1] == 1 {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
			inverted := false
			ans := make([]int, 0)
			j := n - 1
			for i := 0; i < n; i++ {
				for j >= 0 && (inverted && a[j] == 1 || !inverted && a[j] == 0) {
					j--
				}
				if j >= 0 {
					ans = append(ans, j+1)
				} else {
					ans = append(ans, 0)
				}
				inverted = !inverted

			}
			for i := n - 1; i >= 0; i-- {
				write(f, ans[i], " ")
			}
			write(f, "\n")
		}
	}
}
