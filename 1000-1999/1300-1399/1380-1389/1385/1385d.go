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
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		ce := make([][]int, 2*n)
		c := make([][]int, 2*n)
		for i := 0; i < 2*n; i++ {
			ce[i] = make([]int, 27)
			c[i] = make([]int, 27)
		}
		for i := n; i < 2*n; i++ {
			for j := 0; j < 27; j++ {
				if s[i-n] != 'a'+byte(j) {
					ce[i][j]++
					c[i][j]++
				}
			}
		}
		for i := n - 1; i > 0; i-- {
			for j := 0; j < 27; j++ {
				ce[i][j] = ce[2*i][j] + ce[2*i+1][j]
			}
		}
		for i := n - 1; i > 0; i-- {
			for j := 0; j < 26; j++ {
				c[i][j] = min(ce[i*2][j]+c[2*i+1][j+1], c[i*2][j+1]+ce[2*i+1][j])
			}
		}
		// for i := 0; i < 2*n; i++ {
		// 	write(f, ce[i], "\n")
		// }
		write(f, c[1][0], "\n")
	}
}
