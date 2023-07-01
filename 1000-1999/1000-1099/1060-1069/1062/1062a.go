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
	ans := 0
	start, end := 0, 0
	for i := 1; i < n; i++ {
		if a[i] == a[i-1]+1 {
			end++
		} else {
			if end-start >= 1 {
				if start == 0 && a[0] == 1 {
					ans = end - start
				} else {
					ans = max(ans, end-start-1)
				}
			}
			start, end = i, i
		}
	}
	if end-start >= 1 {
		if start == 0 && a[1] == 2 || end == n-1 && a[n-2] == 999 {
			ans = max(ans, end-start)
		} else {
			ans = max(ans, end-start-1)
		}
	}

	write(f, ans, "\n")
}
