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
	ok := true
	c := make([]int, 100001)
	for i := 0; i <= 100000; i++ {
		c[i] = -1
	}
	for i := 0; i < n; i++ {
		var x, k int
		fmt.Fscan(reader, &x, &k)
		if x > c[k]+1 {
			ok = false
		} else {
			c[k] = max(c[k], x)
		}
	}
	if ok {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
