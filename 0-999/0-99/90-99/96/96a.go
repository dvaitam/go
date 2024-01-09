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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	count := 1
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			ans = max(ans, count)
			count = 1
		}
	}
	ans = max(ans, count)
	if ans >= 7 {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
