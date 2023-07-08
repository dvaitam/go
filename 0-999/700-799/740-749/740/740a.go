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
	var n, a, b, c int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &a, &b, &c)
	if n%4 == 0 {
		write(f, "0\n")
	} else if n%4 == 1 {
		write(f, min(3*a, min(b+a, c)), "\n")
	} else if n%4 == 2 {
		write(f, min(min(2*a, b), 2*c), "\n")
	} else {
		write(f, min(min(a, 3*c), b+c), "\n")
	}

}
