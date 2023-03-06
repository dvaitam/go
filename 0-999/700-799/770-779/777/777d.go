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
func get_less(s string, t string) string {
	for i := 0; i < len(s); i++ {
		if i == len(t) {
			return s[:len(t)]
		} else if s[i] > t[i] {
			return s[:i]
		} else if s[i] < t[i] {
			return s
		}
	}
	return s
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]string, n)
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	ans[n-1] = s[n-1]
	for i := n - 2; i >= 0; i-- {
		ans[i] = get_less(s[i], ans[i+1])
	}
	for i := 0; i < n; i++ {
		write(f, ans[i], "\n")
	}
}
