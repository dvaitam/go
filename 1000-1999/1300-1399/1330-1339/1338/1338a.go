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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		i := 0
		j := 0
		diff := int64(0)

		for j < n {
			for j < n && a[j] <= a[i] {
				diff = max(diff, a[i]-a[j])
				j++
			}
			i = j
		}
		if diff == 0 {
			write(f, "0\n")
		} else {

			curr := diff
			x := 0
			for curr > 0 {
				curr -= (1 << x)
				x++
			}

			write(f, x, "\n")
		}
	}
}
