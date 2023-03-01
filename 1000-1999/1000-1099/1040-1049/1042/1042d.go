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
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
func insert_element(a []int64, x int64) []int64 {
	i := bisect_left(a, x)
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	a[i] = x
	return a
}
func main() {
	var n int
	var t int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &t)
	a := make([]int64, n)
	c := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if i == 0 {
			c[i] = a[i]
		} else {
			c[i] = c[i-1] + a[i]
		}
	}
	ans := int64(0)
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	//write(f, c, "\n")
	ans += int64(bisect_left(c, t))
	//write(f, "ans ", ans, "\n")
	//sorted_array := make([]int64, 0)
	cu := int64(0)
	for i := 0; i < n; i++ {
		cu += a[i]
		to := bisect_left(c, cu)
		copy(c[to:], c[to+1:])
		c = c[:len(c)-1]
		//	sorted_array = insert_element(sorted_array, cu)
		//	write(f, sorted_array, "\n")
		t += a[i]
		ans += int64(bisect_left(c, t))
		//write(f, "ans after 1 ", ans, "\n")
		//ans -= int64(bisect_left(sorted_array, t))
		//write(f, "ans after 2 ", ans, "\n")
	}
	write(f, ans, "\n")

}
