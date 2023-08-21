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
		var n0, n1, n2 int
		fmt.Fscan(reader, &n0, &n1, &n2)

		for i := 0; i <= n2 && n2 > 0; i++ {
			write(f, "1")
		}
		for i := 0; i <= n0 && n0 > 0; i++ {
			write(f, "0")
		}
		flip := true

		if n0 > 0 && n2 > 0 {
			n1--
		}
		if n0 > 0 {
			flip = false
		}
		if n0 == 0 && n2 == 0 {
			n1++
		}
		for i := 0; i < n1; i++ {
			if flip {
				write(f, "0")
			} else {
				write(f, "1")
			}
			flip = !flip
		}
		write(f, "\n")
	}
}
