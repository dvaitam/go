// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		fmt.Fscan(reader, &n)
		b := make([]int, n+2)
		for i := 0; i < n+2; i++ {
			fmt.Fscan(reader, &b[i])
		}
		sort.Ints(b)
		ps := 0
		for i := 0; i < n; i++ {
			ps += b[i]
		}
		if ps == b[n] || ps == b[n+1] {
			for i := 0; i < n; i++ {
				write(f, b[i])
				if i == n-1 {
					write(f, "\n")
				} else {
					write(f, " ")
				}
			}
		} else {
			rem := ps - (b[n+1] - b[n])
			ok := false
			skip := -1
			for i := 0; i < n; i++ {
				if b[i] == rem {
					ok = true
					skip = i
					break
				}
			}
			if ok {
				for i := 0; i < n; i++ {
					if i == skip {
						continue
					}
					write(f, b[i], " ")
				}
				write(f, b[n], "\n")
			} else {
				write(f, "-1\n")
			}
		}
	}
}
