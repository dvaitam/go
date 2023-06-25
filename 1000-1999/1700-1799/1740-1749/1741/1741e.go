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
		fmt.Fscan(reader, &n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		p := make([]bool, n)
		for i := 0; i < n; i++ {
			if i == 0 {
				if b[i] < n {
					p[b[i]] = true
				}
			} else {
				if b[i] == i || i-b[i]-1 >= 0 && p[i-b[i]-1] {
					p[i] = true
				}
				if p[i-1] {
					if i+b[i] < n {
						p[i+b[i]] = true
					}
				}
			}
		}
		if p[n-1] {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
