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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		if n%3 == 0 {
			write(f, n/3, n/3, n/3, "\n")
		} else {
			if n%2 == 1 {
				write(f, (n-1)/2, (n-1)/2, 1, "\n")
			} else {
				if n%4 == 0 {
					write(f, n/2, n/4, n/4, "\n")
				} else {
					write(f, 2, (n-2)/2, (n-2)/2, "\n")
				}
			}
		}
	}
}
