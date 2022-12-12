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
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sm := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
		if t[i] == 1 {
			sm += a[i]
		}
	}
	for i := 0; i < k; i++ {
		if t[i] == 0 {
			sm += a[i]
		}
	}
	ans := sm
	for i := k; i < n; i++ {
		if t[i] == 0 {
			sm += a[i]
		}
		if t[i-k] == 0 {
			sm -= a[i-k]
		}
		ans = max(ans, sm)
	}
	write(f, ans, "\n")

}
