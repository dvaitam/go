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
		var l, r int
		fmt.Fscan(reader, &l, &r)
		if (r-l)%2 == 1 {
			if l%2 == 0 {
				write(f, -1*(r-l+1)/2, "\n")
			} else {
				write(f, (r-l+1)/2, "\n")
			}
		} else {
			if l%2 == 1 {
				write(f, ((r-l)/2)-r, "\n")
			} else {
				write(f, l+(r-l)/2, "\n")
			}
		}
	}
}
