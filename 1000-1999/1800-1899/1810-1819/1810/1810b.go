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
		if n%2 == 0 {
			write(f, "-1\n")
		} else {
			ans := make([]int, 0)
			for n > 1 {
				if ((n+1)/2)%2 == 0 {
					ans = append(ans, 2)
					n = (n - 1) / 2
				} else {
					ans = append(ans, 1)
					n = (n + 1) / 2
				}
			}
			write(f, len(ans), "\n")
			for i := len(ans) - 1; i >= 0; i-- {
				write(f, ans[i], " ")
			}
			write(f, "\n")

		}
	}
}
