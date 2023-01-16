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
		var x1, y1, x2, y2 int64
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		r := x2 - x1
		d := y2 - y1
		if r == 0 && d == 0 {
			write(f, "1\n")
		} else if r == 0 || d == 0 {
			write(f, "1\n")
		} else {
			m := min(r, d)
			if m == 1 {
				write(f, r+d, "\n")
			} else {
				ans := (m-1)*m + (r+d-1-2*(m-1))*m + 1
				write(f, ans, "\n")

			}

		}
	}
}
