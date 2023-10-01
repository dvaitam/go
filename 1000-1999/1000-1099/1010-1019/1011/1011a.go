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
	var n, k int
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k, &s)
	c := make([]bool, 26)
	for i := 0; i < n; i++ {
		c[s[i]-'a'] = true
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if k > 0 {
			if c[i] {
				ans += 1 + i
				i++
				k--

			}
		}
	}
	if k == 0 {
		write(f, ans, "\n")
	} else {
		write(f, "-1\n")
	}
}
