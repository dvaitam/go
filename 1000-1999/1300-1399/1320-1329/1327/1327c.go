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
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &k)
	for i := 0; i < 2*k; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
	}
	ans := make([]byte, 0)
	for i := 0; i < n-1; i++ {
		ans = append(ans, 'U')
	}
	for i := 0; i < m-1; i++ {
		ans = append(ans, 'L')
	}
	direction := true
	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			if direction {
				ans = append(ans, 'R')
			} else {
				ans = append(ans, 'L')
			}
		}
		ans = append(ans, 'D')
		direction = !direction
	}
	write(f, len(ans), "\n", string(ans), "\n")
}
