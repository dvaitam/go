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
	var n, q int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	start := make([]int, n+1)
	end := make([]int, n+1)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		start[l]++
		end[r]++
	}
	times := make([]int, n+1)
	curr := 0
	for i := 1; i <= n; i++ {
		curr += start[i]
		times[i] = curr
		curr -= end[i]
	}
	sort.Ints(times)
	sort.Ints(a)
	ans := int64(0)
	for i := 1; i <= n; i++ {
		ans += int64(times[i] * a[i])
	}
	write(f, ans, "\n")
}
