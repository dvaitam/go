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
		var s string
		fmt.Fscan(reader, &n, &s)
		ops := 0
		i := 0
		start := -1
		for i < n {
			if start != -1 && s[i] == ')' {
				ops++
				i++
				start = -1
			} else if start != -1 {
				i++
			} else {
				if i < n-1 {
					if s[i] == '(' || s[i+1] == ')' {
						ops++
						i += 2
					} else {
						start = i
						i++
					}
				} else {
					start = n - 1
					i++
				}
			}
		}
		if start == -1 {
			write(f, ops, 0, "\n")
		} else {
			write(f, ops, n-start, "\n")
		}
	}
}
