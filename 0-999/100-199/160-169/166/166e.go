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
	a := make([]int64, n+1)
	b := make([]int64, n+1)
	b[1] = 1
	for i := 2; i <= n; i++ {
		a[i] = (3 * b[i-1]) % 1000000007
		if i%2 == 0 {
			b[i] = (3*b[i-1] - 1) % 1000000007
		} else {
			b[i] = (3*b[i-1] + 1) % 1000000007
		}
	}
	write(f, a[n], "\n")
}
