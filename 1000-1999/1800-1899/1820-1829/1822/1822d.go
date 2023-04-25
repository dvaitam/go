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
		fmt.Fscan(reader, &n)
		if n == 1 {
			write(f, "1\n")
		} else if n%2 == 1 {
			write(f, "-1\n")
		} else {
			if n%8 == 6 {
				write(f, n, " ")
				neg := true
				for i := 1; i < n/2; i++ {
					if neg {
						write(f, n-i, " ")
					} else {
						write(f, i, " ")
					}
					neg = !neg
				}
				write(f, n/2, " ")
				for i := (n / 2) - 1; i > 0; i-- {
					if neg {
						write(f, n-i, " ")
					} else {
						write(f, i, " ")
					}
					neg = !neg
				}
				write(f, "\n")

			} else {
				write(f, n, " ")
				neg := false
				for i := 1; i < n/2; i++ {
					if neg {
						write(f, n-i, " ")
					} else {
						write(f, i, " ")
					}
					neg = !neg
				}
				write(f, n/2, " ")
				for i := (n / 2) - 1; i > 0; i-- {
					if neg {
						write(f, n-i, " ")
					} else {
						write(f, i, " ")
					}
					neg = !neg
				}
				write(f, "\n")

			}
		}
	}
}
