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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	mb := make([][]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		mb[i] = make([]bool, m)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mb[i][j] = true
		}
	}
	ans := 0
	for bi := 8; bi >= 0; bi-- {
		curr := 1 << bi
		min_count := 1000000
		for i := 0; i < n; i++ {
			row_count := 0
			for j := 0; j < m; j++ {
				if mb[i][j] && (a[i]&b[j]&curr != curr) {
					row_count++
				}
			}
			min_count = min(min_count, row_count)
		}
		if min_count == 0 {
			ans = ans | curr
		} else {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if mb[i][j] && (a[i]&b[j]&curr == curr) {
						mb[i][j] = false
					}
				}
			}
		}
	}
	write(f, ans, "\n")

}
