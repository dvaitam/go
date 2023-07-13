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
	m := map[int]bool{}
	cap := 0
	curr := 0
	for i := 0; i < n; i++ {
		var s string
		var x int
		fmt.Fscan(reader, &s, &x)
		if s[0] == '+' {
			m[x] = true
			curr++
			cap = max(cap, curr)
		} else {
			if m[x] {
				curr--
			} else {
				cap++
			}
			m[x] = false
		}
	}
	write(f, cap, "\n")
}
