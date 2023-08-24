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
	var n int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	ans := n
	for i := n - 1; i > 0; i-- {
		ans *= i
		ans = ans % 1000000007
	}
	sub := int64(1)
	for i := int64(0); i < n-1; i++ {
		sub = sub * 2
		sub = sub % 1000000007
	}
	ans = (ans - sub + 1000000007) % 1000000007
	write(f, ans, "\n")
}
