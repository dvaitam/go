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
	var n, m, q int
	var s, t string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &q, &s, &t)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		ok := true
		for j := i; j < i+m; j++ {
			if j >= n {
				ok = false
				break
			}
			if s[j] != t[j-i] {
				ok = false
				break
			}
		}
		if ok {
			a[i] = 1
		}
	}
	for i := 1; i < n; i++ {
		a[i] += a[i-1]
	}
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		l--
		r--
		if r-l+1 < m {
			write(f, 0, "\n")
		} else {
			r -= m - 1
			ans := a[r]
			if l > 0 {
				ans -= a[l-1]
			}
			write(f, ans, "\n")
		}

	}
}
