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
		iden := true
		first, last := n, -1
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] != i+1 {
				iden = false
				first = min(first, i)
				last = max(last, i)
			}
		}
		if iden {
			write(f, "0\n")
		} else {
			extra := false
			for i := first; i <= last; i++ {
				if a[i] == i+1 {
					extra = true
				}
			}
			if extra {
				write(f, "2\n")
			} else {
				write(f, "1\n")
			}
		}
	}
}
