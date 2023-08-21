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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	var s, t string
	fmt.Fscan(reader, &s, &t)
	a := make([]int, m)
	b := make([]int, m)
	j := 0
	for i := 0; i < n; i++ {
		if j < m && s[i] == t[j] {
			a[j] = i
			j++
		}
	}
	j = m - 1
	for i := n - 1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			b[j] = i
			j--
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		if i < m-1 {
			ans = max(ans, b[i+1]-a[i])
		}
	}
	write(f, ans, "\n")
}
