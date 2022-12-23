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
		mi := -1
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i > 0 {
				if mi == -1 {
					mi = a[i]
				} else {
					mi = min(mi, a[i])
				}
			}
		}
		if a[0] > mi {
			// fix
			write(f, "Alice\n")
		} else {
			write(f, "Bob\n")
		}

	}
}
