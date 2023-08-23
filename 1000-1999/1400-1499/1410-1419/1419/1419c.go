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
		var n, x int
		fmt.Fscan(reader, &n, &x)
		a := make([]int, n)
		delta := 0
		all_equal := true
		count := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] != x {
				all_equal = false
			}
			if a[i] == x {
				count++
			}
			delta += x - a[i]
		}

		if all_equal {
			write(f, "0\n")
		} else {
			if delta == 0 || count >= 1 {
				write(f, "1\n")
			} else {
				write(f, "2\n")
			}
		}
		// if t == 4 {
		// 	write(f, n, x, count, a)
		// 	write(f, "delta ", delta, "\n")
		// }

	}
}
