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
		var n, l, r int64
		fmt.Fscan(reader, &n, &l, &r)
		if l == n*(n-1)+1 {
			write(f, 1, "\n")
		} else {
			start := int64(1)
			sm := int64(0)
			for sm+2*(n-start) < l {
				sm += 2 * (n - start)
				start++
			}
			second := start + (l-sm)/2
			if (l-sm)%2 == 0 {
				write(f, second, " ")
				second++
				l++
			} else {
				second++
			}
			for l <= r {
				if second == n+1 {
					start++
					second = start + 1
				}
				if start >= n {
					write(f, 1)
					break
				}
				if r-l > 0 {
					write(f, start, " ", second, " ")
					second++
					l += 2
				} else {
					write(f, start, " ")
					l += 1
				}
				//write(f, "second here is ", second, "\n")
				if second == n+1 {
					start++
					second = start + 1
				}

			}
			write(f, "\n")
		}

	}
}
