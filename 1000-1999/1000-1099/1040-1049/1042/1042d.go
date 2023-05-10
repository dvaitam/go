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
func bisect_left(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
}
func lb(x int) int {
	return x & (-x)
}
func mod(bit []int, x int) {
	for i := x; i < len(bit); i += lb(i) {
		bit[i]++
	}
}
func query(bit []int, x int) int {
	r := 0
	for i := x; i > 0; i -= lb(i) {
		r += bit[i]
	}
	return r
}

func main() {
	var n int
	var t int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &t)
	a := make([]int64, n)
	c := make([]int64, n+1)
	uniq := map[int64]bool{}
	li := make([]int64, 0)
	bit := make([]int, 402020)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		c[i+1] = c[i] + a[i]
		if !uniq[c[i+1]] {
			li = append(li, c[i+1])
			uniq[c[i+1]] = true
		}
		if !uniq[c[i+1]-t] {
			li = append(li, c[i+1]-t)
			uniq[c[i+1]-t] = true
		}
	}
	if uniq[0] {
		li = append(li, 0)
	}
	sort.Slice(li, func(i, j int) bool { return li[i] < li[j] })
	ans := int64(0)
	for i := 1; i <= n; i++ {
		lb := bisect_left(li, c[i-1]) + 1
		mod(bit, lb)
		llb := bisect_left(li, c[i]-t) + 1
		ans += int64(i - query(bit, llb))
	}

	write(f, ans, "\n")

}
