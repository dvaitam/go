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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &s)
	h := make([]int, n)

	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			if s[i] == '[' {
				h[i] = h[i-1] - 2
			} else {
				h[i] = h[i-1] + 2
			}
		} else {
			h[i] = h[i-1]
		}
	}
	mi := 100000
	for i := 0; i < n; i++ {
		mi = min(mi, h[i])
	}
	for i := 0; i < n; i++ {
		h[i] -= mi
		h[i] += 3
	}
	mx := 0
	for i := 0; i < n; i++ {
		mx = max(mx, h[i])
	}
	ans := make([][]byte, mx)
	for i := 0; i < mx; i++ {
		ans[i] = make([]byte, 3*n)
	}
	for i := 0; i < mx; i++ {
		for j := 0; j < 3*n; j++ {
			ans[i][j] = ' '
		}
	}
	curr := 0
	for i := 0; i < n; i++ {
		for j := (mx / 2) - h[i]/2; j <= (mx/2)+h[i]/2; j++ {
			if j == (mx/2)-h[i]/2 || j == (mx/2)+h[i]/2 {
				ans[j][curr] = '+'
				if s[i] == '[' {
					ans[j][curr+1] = '-'
				} else {
					ans[j][curr-1] = '-'
				}
			} else {
				ans[j][curr] = '|'
			}
		}
		if s[i] == '[' && s[i+1] == ']' {
			curr += 4
		} else {
			curr += 1
		}
	}
	for i := 0; i < mx; i++ {
		write(f, string(ans[i][:curr]), "\n")
	}

}
