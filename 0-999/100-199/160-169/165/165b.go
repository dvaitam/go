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
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		mitd := mid
		sm := 0
		for mid > 0 {
			sm += mid
			mid = mid / k
		}
		write(f, left, right, mitd, sm, "\n")
		if sm >= n {
			right = mitd
		} else {
			left = mitd + 1
		}
	}
	write(f, left, "\n")
}
