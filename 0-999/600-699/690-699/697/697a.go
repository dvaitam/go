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
	var t, x, s int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &t, &s, &x)
	if x < t {
		write(f, "NO\n")
	} else {
		rem := (x - t) % s
		if x != t+1 && (rem == 0 || rem == 1) {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
