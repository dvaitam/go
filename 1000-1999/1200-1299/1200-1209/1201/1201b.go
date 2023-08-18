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
	a := make([]int64, n)
	sm := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		sm += a[i]
	}
	ok := true
	for i := 0; i < n; i++ {
		if a[i]*2 > sm {
			ok = false
			break
		}
	}
	if ok && sm%2 == 0 {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
