// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	//write(f, a, "\n")
	ans := a[n-1] - a[0]
	for i := n; i < m; i++ {
		ans = min(ans, a[i]-a[i-n+1])
	}
	write(f, ans, "\n")
}
