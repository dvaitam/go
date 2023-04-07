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
func get_wins(a1 int, a2 int, a3 int, b1 int, b2 int, b3 int) int {
	return min(b1, a2) + min(b2, a3) + min(b3, a1)
}
func get_losses(a1 int, a2 int, a3 int, b1 int, b2 int, b3 int) int {
	return max(b1-min(b1, a1+a2), max(b2-min(b2, a2+a3), b3-min(b3, a1+a3)))
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	var a1, a2, a3, b1, b2, b3 int
	fmt.Fscan(reader, &a1, &a2, &a3, &b1, &b2, &b3)
	write(f, get_losses(a1, a2, a3, b1, b2, b3), get_wins(b1, b2, b3, a1, a2, a3), "\n")

}
