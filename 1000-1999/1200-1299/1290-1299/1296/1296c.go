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
		m := map[int64]int{}
		mx := n
		l, r := -1, -1
		x, y := int64(0), int64(0)
		for i := 0; i < n; i++ {
			if s[i] == 'R' {
				x++
			} else if s[i] == 'L' {
				x--
			} else if s[i] == 'U' {
				y++
			} else {
				y--
			}
			co := x*1000000000 + y
			if m[co] > 0 || (x == 0 && y == 0) {
				if m[co] > 0 {
					if i+1-m[co] < mx {
						mx = i + 1 - m[co]
						l, r = m[co]+1, i+1
					}
				} else {
					if i < mx {
						mx = i
						l, r = 1, i+1
					}
				}

			}
			m[co] = i + 1
		}

		if l == -1 {
			if x == 0 && y == 0 {
				write(f, 1, n, "\n")
			} else {
				write(f, "-1\n")
			}
		} else {
			write(f, l, r, "\n")
		}
	}
}
