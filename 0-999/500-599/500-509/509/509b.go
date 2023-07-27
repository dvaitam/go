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
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, 0)
	}
	ok := true
	for c := 1; c <= k; c++ {
		mi := 101
		for i := 0; i < n; i++ {
			mi = min(mi, a[i]-len(ans[i]))

		}
		if mi == 101 {
			break
		}
		for i := 0; i < n; i++ {
			if a[i]-len(ans[i]) > mi {
				for j := 0; j < mi+1; j++ {
					ans[i] = append(ans[i], c)
				}
			} else {
				for j := 0; j < mi; j++ {
					ans[i] = append(ans[i], c)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if a[i] != len(ans[i]) {
			ok = false
			break
		}
	}
	if ok {
		write(f, "YES\n")
		for i := 0; i < n; i++ {
			for j := 0; j < len(ans[i]); j++ {
				write(f, ans[i][j], " ")
			}
			write(f, "\n")
		}
	} else {
		write(f, "NO\n")
	}
}
