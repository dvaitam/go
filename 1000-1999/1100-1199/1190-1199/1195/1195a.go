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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	c := make([]int, k+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		c[a[i]]++
	}
	ans := 0
	rem := 0
	for i := 1; i <= k; i++ {
		ans += c[i] - (c[i] % 2)
		rem += c[i] % 2
	}
	ans += rem / 2
	ans += rem % 2
	write(f, ans, "\n")
}
