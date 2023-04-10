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
		first1, first2 := 2*n-1, 1
		for i := 0; i < n; i += 2 {
			write(f, first1, first2, " ")
			first1 -= 2
			first2 += 2
		}
		write(f, "\n")
		first1, first2 = 2, n+2
		for i := 0; i < n; i += 2 {
			write(f, first1, first2, " ")
			first1 += 2
			first2 += 2
		}
		write(f, "\n")

	}
}
