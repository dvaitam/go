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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		mx := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			mx = max(mx, a[i])
		}
		start := 0
		end := n - 1
		if mx == a[start] {
			start++
			for k := end; k > 0; k-- {
				write(f, a[k], " ")
			}
			write(f, a[0], "\n")

		} else if mx == a[end] {
			write(f, a[end])
			end--
			for k := end; k >= 0; k-- {
				write(f, " ", a[k])
			}
			write(f, "\n")
		} else {
			write(f, "-1\n")
		}

	}
}
