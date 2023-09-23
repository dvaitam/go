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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int64, n)
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	tc := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])

		tc += a[i]
	}
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	if tc <= b[0]+b[1] {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
