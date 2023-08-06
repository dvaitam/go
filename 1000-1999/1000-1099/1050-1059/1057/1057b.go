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
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &r[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		curr := 0
		for j := i; j < n; j++ {
			curr += r[j]
			if curr > 100*(j-i+1) {
				ans = max(ans, j-i+1)
			}
		}
	}
	write(f, ans, "\n")

}
