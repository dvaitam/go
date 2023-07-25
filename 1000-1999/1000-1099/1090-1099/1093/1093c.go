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
	b := make([]int64, n/2)
	for i := 0; i < n/2; i++ {
		fmt.Fscan(reader, &b[i])
	}
	a := make([]int64, n)
	a[0] = 0
	a[n-1] = b[0]
	for i := 1; i < n/2; i++ {
		right := n - i - 1
		if b[i]-a[i-1] <= a[right+1] {
			a[i] = a[i-1]
			a[right] = b[i] - a[i-1]
		} else {
			a[right] = a[right+1]
			a[i] = b[i] - a[right]
		}
	}
	for i := 0; i < n; i++ {
		write(f, a[i], " ")
	}
	write(f, "\n")
}
