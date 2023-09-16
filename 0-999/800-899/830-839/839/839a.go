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
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	ans := -1
	rem := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		rem += a[i]
		give := min(8, rem)
		give = min(k, give)
		if k > 0 {
			k -= give
			rem -= give
			if k == 0 {
				ans = i + 1
				break
			}
		}
	}
	write(f, ans, "\n")

}
