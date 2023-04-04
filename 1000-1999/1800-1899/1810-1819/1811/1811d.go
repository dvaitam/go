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
		var x, y int64
		fmt.Fscan(reader, &n, &x, &y)
		a := make([]int64, n+1)
		a[0] = 1
		a[1] = 1
		for i := 2; i <= n; i++ {
			a[i] = a[i-1] + a[i-2]
		}
		//write(f, a, "\n")
		isx := true
		for i := n; i > 1; i-- {
			if isx {
				if y > a[i] {
					y -= a[i]
				}
			} else {
				if x > a[i] {
					x -= a[i]
				}
			}
			isx = !isx
		}
		//	write(f, x, y, "\n")
		if n%2 == 1 {
			if x == 1 && y == 1 || x == 1 && y == 2 {
				write(f, "YES\n")
			} else {
				write(f, "NO\n")
			}
		} else {
			if x == 1 && y == 1 || x == 2 && y == 1 {
				write(f, "YES\n")
			} else {
				write(f, "NO\n")
			}
		}
	}
}
