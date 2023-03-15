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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	o := make([]int, n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		add := 0
		if a[i] == 1 {
			add = 1
		}
		if i == 0 {
			o[i] = add
		} else {
			o[i] = o[i-1] + add
		}
	}
	for i := n - 1; i >= 0; i-- {
		add := 0
		if a[i] == 2 {
			add = 1
		}
		if i == n-1 {
			t[i] = add
		} else {
			t[i] = t[i+1] + add
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		twos := 0
		ones := 0
		for j := i; j < n; j++ {
			if a[j] == 2 {
				twos++
			} else {
				ones = max(twos+1, ones+1)
			}
			dp[i][j] = max(ones, twos)
		}
	}
	//write(f, dp, "\n")
	ans := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			tmp := 0
			if i > 0 {
				tmp += o[i-1]
			}
			if j < n-1 {
				tmp += t[j+1]
			}
			tmp += dp[i][j]
			ans = max(ans, tmp)
		}
	}
	write(f, ans, "\n")

}
