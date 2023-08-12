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
	p := make([]string, n)
	s := make([]int, n)
	m := map[string]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i], &s[i])
		m[p[i]] += s[i]
	}
	mx := 0
	for _, score := range m {
		mx = max(mx, score)
	}
	// if n == 50 {
	// 	write(f, m, "\n")
	// }
	mn := map[string]int{}
	ans := ""
	for i := 0; i < n; i++ {
		mn[p[i]] += s[i]
		if mn[p[i]] >= mx && m[p[i]] == mx {
			ans = p[i]
			break
		}
	}
	write(f, ans, "\n")
}
