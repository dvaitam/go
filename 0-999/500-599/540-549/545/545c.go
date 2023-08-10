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
	x := make([]int, n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &h[i])
	}
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			l[i] = 1
			if i == n-1 || x[i+1]-x[i] > h[i] {
				r[i] = 1
			}
		} else {
			val1 := l[i-1]
			if x[i]-x[i-1] > h[i] {
				val1++
			}
			val2 := r[i-1]
			if x[i]-x[i-1]-h[i-1] > h[i] {
				val2++
			}
			l[i] = max(val1, val2)
			val3 := max(l[i-1], r[i-1])
			if i+1 == n || x[i+1]-x[i] > h[i] {
				val3++
			}
			r[i] = val3
		}
	}
	//write(f, l, "\n", r, "\n")
	ans := max(l[0], r[0])
	for i := 1; i < n; i++ {
		ans = max(ans, l[i])
		ans = max(ans, r[i])
	}
	write(f, ans, "\n")
}
