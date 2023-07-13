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
	for i := 1; i <= 1200; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if i <= b[j] {
				if s[j] == '1' {
					count++
				}
			} else {
				times := (i - b[j]) / a[j]
				if times%2 == 0 {
					if s[j] == '0' {
						count++
					}
				} else {
					if s[j] == '1' {
						count++
					}
				}
			}
		}
		ans = max(ans, count)
	}
	write(f, ans, "\n")
}
