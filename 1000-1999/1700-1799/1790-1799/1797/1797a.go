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
func get_min(n int, m int, x int, y int) int {
	if (x == 1 || x == n) && (y == 1 || y == m) {
		return 2
	}
	if x == 1 || y == 1 || x == n || y == m {
		return 3
	}
	return 4
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		var x1, y1, x2, y2 int
		fmt.Fscan(reader, &n, &m, &x1, &y1, &x2, &y2)
		write(f, min(get_min(n, m, x1, y1), get_min(n, m, x2, y2)), "\n")
	}
}
