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
		a := make([]int, n)
		m := map[int]int{}
		ans := -1
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if m[a[i]] > 0 {
				if ans == -1 {
					ans = i + 2 - m[a[i]]
				} else {
					ans = min(ans, i+2-m[a[i]])
				}
			}
			m[a[i]] = i + 1
		}
		write(f, ans, "\n")

	}
}
