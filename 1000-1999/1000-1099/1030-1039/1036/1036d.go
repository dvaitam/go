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
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	b := make([]int64, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	ca := make([]int64, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ca[i] = a[i]
		} else {
			ca[i] = ca[i-1] + a[i]
		}
	}
	cb := make([]int64, m)
	for i := 0; i < m; i++ {
		if i == 0 {
			cb[i] = b[i]
		} else {
			cb[i] = cb[i-1] + b[i]
		}
	}
	if ca[n-1] != cb[m-1] {
		write(f, "-1\n")
	} else {
		ans := 0
		i, j := 0, 0
		for i < n && j < m {
			for i < n && ca[i] < cb[j] {
				i++
			}
			for j < m && cb[j] < ca[i] {
				j++
			}
			if i < n && j < m && ca[i] == cb[j] {
				ans++
				i++
			}
		}
		write(f, ans, "\n")
	}
}
