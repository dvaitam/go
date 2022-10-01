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
		var n, m, sx, sy, d int
		fmt.Fscan(reader, &n, &m, &sx, &sy, &d)
		if (sx-d <= 1 && sy-d <= 1) || (sx+d >= n && sy+d >= m) || (sx-d <= 1 && sx+d >= n) || (sy-d <= 1 && sy+d >= m) {
			write(f, -1, "\n")
		} else {
			write(f, n+m-2, "\n")
		}
	}
}
