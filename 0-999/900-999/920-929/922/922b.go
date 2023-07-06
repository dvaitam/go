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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	ans := 0
	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			c := a ^ b
			if c < a || c < b || c > n {
				continue
			}
			if a+b <= c || a+c <= b || b+c <= a {
				continue
			}
			//write(f, a, b, c, "\n")
			ans++

		}
	}
	write(f, ans, "\n")
}
