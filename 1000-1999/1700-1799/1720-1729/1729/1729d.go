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
		x := make([]int, n)
		y := make([]int, n)
		d := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &x[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &y[i])
			d[i] = y[i] - x[i]
		}
		sort.Ints(d)
		start := 0
		end := n - 1
		ans := 0
		for start < end {
			if d[start]+d[end] >= 0 {
				start++
				end--
				ans++
			} else {
				start++
			}
		}
		write(f, ans, "\n")
	}
}
