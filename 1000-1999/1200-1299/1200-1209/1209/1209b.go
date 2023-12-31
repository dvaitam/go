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
	var s string
	fmt.Fscan(reader, &s)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}
	ans := 0
	o := make([][]bool, n)
	m := 1200
	for i := 0; i < n; i++ {
		o[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			o[i][0] = true
		}
		for j := 1; j < m; j++ {
			if j < b[i] {
				o[i][j] = o[i][j-1]
			} else if (b[i]-j)%a[i] == 0 {
				o[i][j] = !o[i][j-1]
			} else {
				o[i][j] = o[i][j-1]
			}
		}
	}
	for j := 0; j < m; j++ {
		count := 0
		for i := 0; i < n; i++ {
			if o[i][j] {
				count++
			}
		}
		ans = max(ans, count)
	}

	write(f, ans, "\n")
}
