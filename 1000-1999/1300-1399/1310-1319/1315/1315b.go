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
		var a, b, p int
		var s string
		fmt.Fscan(reader, &a, &b, &p, &s)
		n := len(s)
		if s[n-2] == 'B' && b > p || s[n-2] == 'A' && a > p {
			write(f, n, "\n")
		} else {
			i := n - 2
			curr := s[n-2]
			for i >= 0 {
				if curr == 'A' && p >= a || curr == 'B' && p >= b {
					if curr == 'A' {
						p -= a
					} else {
						p -= b
					}
					for i >= 0 && s[i] == curr {
						i--
					}
					if i >= 0 {
						curr = s[i]
					}
				} else {
					break
				}
			}
			write(f, i+2, "\n")
		}
	}
}
