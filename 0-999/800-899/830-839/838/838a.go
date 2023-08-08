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
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		var s string
		fmt.Fscan(reader, &s)
		for j := 0; j < m; j++ {
			if s[j] == '1' {
				a[i][j] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				if j > 0 {
					a[i][j] += a[i][j-1]
				}
			} else if j == 0 {
				if i > 0 {
					a[i][j] += a[i-1][j]
				}
			} else {
				a[i][j] += a[i][j-1] + a[i-1][j] - a[i-1][j-1]
			}
		}
	}
	//write(f, a, "\n")
	ans := m * n
	for k := 2; k <= min(n, m); k++ {
		kn, km := n, m
		if kn%k != 0 {
			kn += k - kn%k
		}
		if km%k != 0 {
			km += k - km%k
		}
		count := 0
		for i := k - 1; i < kn; i += k {
			for j := k - 1; j < km; j += k {
				in, jm := min(i, n-1), min(j, m-1)
				box := a[in][jm]
				if i-k >= 0 {
					box -= a[i-k][jm]
				}
				if j-k >= 0 {
					box -= a[in][j-k]
				}
				if i-k >= 0 && j-k >= 0 {
					box += a[i-k][j-k]
				}
				count += min(box, k*k-box)
			}
		}
		ans = min(count, ans)
	}
	write(f, ans, "\n")
}
