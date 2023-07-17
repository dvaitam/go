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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	u := make([]int, 0)
	for i := 0; i < n; i++ {
		if i == 0 {
			u = append(u, a[i])
		} else {
			if a[i] != a[i-1] {
				u = append(u, a[i])
			}
		}
	}
	for i := 0; i < k; i++ {
		if i < len(u) {
			if i == 0 {
				write(f, u[i], "\n")
			} else {
				write(f, u[i]-u[i-1], "\n")
			}
		} else {
			write(f, "0\n")
		}
	}
}
