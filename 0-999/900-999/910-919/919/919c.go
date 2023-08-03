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
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &k)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if s[i][j] == '.' {
				count++
			} else {
				if count >= k {
					ans += count - k + 1
				}
				count = 0
			}
		}
		if count >= k {
			ans += count - k + 1
		}
	}
	for i := 0; i < m; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if s[j][i] == '.' {
				count++
			} else {
				if count >= k {
					ans += count - k + 1
				}
				count = 0
			}
		}
		if count >= k {
			ans += count - k + 1
		}
	}
	if k == 1 {
		write(f, ans/2, "\n")
	} else {
		write(f, ans, "\n")

	}
}
