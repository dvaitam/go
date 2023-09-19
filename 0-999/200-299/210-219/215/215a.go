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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscan(reader, &m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	ans := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if b[i]%a[j] == 0 {
				ans = max(ans, b[i]/a[j])
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if b[i]%a[j] == 0 && b[i]/a[j] == ans {
				count++
			}
		}
	}
	write(f, count, "\n")

}
