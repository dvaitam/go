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
		if n%2 == 0 {
			pos := k
			if pos%n == 0 {
				pos = n
			} else {
				pos = pos % n
			}
			write(f, pos, "\n")
		} else {
			half := (n / 2)
			rounds := (k - 1) / half
			pos := 1 + (half+1)*rounds + (k-1)%half
			if pos%n == 0 {
				pos = n
			} else {
				pos = pos % n
			}
			write(f, pos, "\n")
		}
	}
}
