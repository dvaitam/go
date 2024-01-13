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
	a := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		total += a[i]
	}
	sort.Ints(a)
	ans := 0
	sm := 0
	for i := n - 1; i >= 0; i-- {
		sm += a[i]
		ans++
		if sm > total-sm {
			break
		}
	}
	write(f, ans, "\n")
}
