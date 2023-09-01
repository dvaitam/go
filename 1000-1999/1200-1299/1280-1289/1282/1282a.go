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
		var a, b, c, r int
		fmt.Fscan(reader, &a, &b, &c, &r)
		if a > b {
			a, b = b, a
		}
		d, e := c-r, c+r
		mi, mx := max(d, a), min(e, b)
		if mx < mi {
			write(f, b-a, "\n")
		} else {
			write(f, b-a-(mx-mi), "\n")
		}
	}
}
