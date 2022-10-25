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
func main() {
	var n, c int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &c)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}
	ans := 0
	for i := 1; i < n; i++ {
		if x[i-1] > x[i]+c {
			ans = max(ans, x[i-1]-x[i]-c)
		}
	}
	write(f, ans, "\n")
}
