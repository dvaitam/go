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
		xor := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			xor = xor ^ a[i]
		}
		if xor == 0 {
			write(f, "0\n")
		} else {
			if n%2 == 0 {
				write(f, "-1\n")
			} else {
				write(f, xor, "\n")
			}
		}
	}
}
