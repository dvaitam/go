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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	min_factor := 1000000000
	for i := 0; i < n; i++ {
		min_factor = min(min_factor, a[i]/n)
	}
	for i := 0; i < n; i++ {
		a[i] = a[i] - n*min_factor
	}

	ans := -1

	for i := 0; i < n; i++ {
		if a[i] <= i {
			ans = i + 1
			break
		}
	}
	if ans == -1 {
		for i := 0; i < n; i++ {
			if a[i] <= n {
				a[i] = 0
			} else {
				a[i] -= n
			}
		}
		for i := 0; i < n; i++ {
			if a[i] <= i {
				ans = i + 1
				break
			}
		}
	}

	write(f, ans, "\n")
}
