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
	var n, q int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	dp := make([]int, n)
	sdp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = 1
		} else if i == n-1 {
			dp[i] = 1
		} else {
			if a[i] <= a[i-1] && a[i] >= a[i+1] {
				dp[i] = 0
			} else {
				dp[i] = 1
			}
		}
		if i == 0 {
			sdp[i] = dp[i]
		} else {
			sdp[i] = sdp[i-1] + dp[i]
		}
	}
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		l--
		r--
		ans := sdp[r]
		if dp[r] == 0 {
			ans++
		}
		if l != r && dp[l] == 0 {
			ans++
		}
		if l > 0 {
			ans -= sdp[l-1]
		}
		write(f, ans, "\n")
	}
}
