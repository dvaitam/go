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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}
	ans := 0

	for i := 0; i < n; i++ {
		if i == 0 {
			ans += h[i] + 1
		} else {
			if h[i] >= h[i-1] {
				ans += 2 + (h[i] - h[i-1])
			} else {
				ans += 2 + (h[i-1] - h[i])
			}
		}
	}
	write(f, ans, "\n")

}
