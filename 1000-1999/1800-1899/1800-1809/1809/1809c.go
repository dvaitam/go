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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		nxt := 1
		ans := make([]int, 0)
		for k-nxt >= 0 {
			ans = append(ans, 2)
			k = k - nxt
			nxt++
		}
		if k > 0 {
			ans = append(ans, -((len(ans)-k)*2 + 1))
		}
		rem := -len(ans)*2 - 1
		for i := len(ans); i < n; i++ {
			ans = append(ans, rem)
		}
		for i := 0; i < n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
	}
}
