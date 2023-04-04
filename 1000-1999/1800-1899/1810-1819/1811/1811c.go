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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		b := make([]int, n-1)
		for i := 0; i < n-1; i++ {
			fmt.Fscan(reader, &b[i])
		}
		ans := make([]int, n)
		for i := 1; i < n-1; i++ {
			ans[i] = min(b[i], b[i-1])
		}
		ans[0] = max(ans[1], b[0])
		ans[n-1] = max(ans[n-2], b[n-2])
		for i := 0; i < n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
	}
}
