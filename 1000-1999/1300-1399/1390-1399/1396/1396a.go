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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	write(f, 1, 1, "\n")
	write(f, 0-a[0], "\n")
	if n > 1 {
		write(f, 2, n, "\n")
		for i := 1; i < n; i++ {
			write(f, a[i]*int64(n-1))
			if i == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
		write(f, 1, n, "\n")
		write(f, 0, " ")
		for i := 1; i < n; i++ {
			write(f, a[i]*int64(-n))
			if i == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	} else {
		write(f, 1, 1, "\n")
		write(f, 0, "\n")
		write(f, 1, 1, "\n")
		write(f, 0, "\n")
	}
}
