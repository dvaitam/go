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
	var s string
	fmt.Fscan(reader, &s)
	ok := false
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			ok = true
			l = i
			r = i + 1
			break
		}
	}
	if ok {
		write(f, "Yes ", l, r, "\n")
	} else {
		write(f, "No\n")
	}
}
