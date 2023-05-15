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

type Pair struct {
	men, key int
	t        int64
}

func time(a int64, b int64, p int64) int64 {
	return abs(a-b) + abs(p-b)
}
func main() {
	var n, k int
	var p int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k, &p)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int64, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &b[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	ans := int64(2000000000)
	for delta := 0; delta <= k-n; delta++ {
		cur := int64(0)
		for i := 0; i < n; i++ {
			cur = max(cur, abs(a[i]-b[i+delta])+abs(b[i+delta]-p))
		}
		ans = min(ans, cur)
	}
	write(f, ans, "\n")
}
