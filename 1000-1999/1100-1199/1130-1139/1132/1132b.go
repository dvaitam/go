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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	total := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		total += int64(a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Fscan(reader, &m)
	q := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &q[i])
		write(f, total-int64(a[q[i]-1]), "\n")
	}

}
