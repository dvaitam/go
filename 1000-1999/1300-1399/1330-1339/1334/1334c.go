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
		a := make([]int64, n)
		b := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i], &b[i])
		}
		mi := a[0]
		ans := int64(0)
		for i := 0; i < n; i++ {
			nxt := i + 1
			if nxt == n {
				nxt = 0
			}
			if a[nxt] > b[i] {
				ans += a[nxt] - b[i]
				a[nxt] = b[i]
			}
			mi = min(mi, a[nxt])

		}
		write(f, ans+mi, "\n")
	}
}
