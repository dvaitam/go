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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s int64
		fmt.Fscan(reader, &n, &s)
		a := make([]int64, n+1)
		sum := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &a[i])
			sum[i] = sum[i-1] + a[i]
		}
		top := 0
		st := make([]int, n+1)
		st[0] = n + 1
		ansl, ansr := 1, 0
		for i := n; i >= 1; i-- {
			for top > 0 && sum[st[top]] >= sum[i] {
				top--
			}
			top++
			st[top] = i
			if a[i]+s < 0 {
				continue
			}
			l, r, ans := 1, top, i
			for l <= r {
				mid := (l + r) >> 1
				if sum[st[mid]]-sum[i-1]+s >= 0 {
					r = mid - 1
					ans = st[mid-1] - 1
				} else {
					l = mid + 1
				}
			}
			if ans-i > ansr-ansl {
				ansl = i
				ansr = ans
			}
		}
		if ansr == 0 {
			write(f, "-1\n")
		} else {
			write(f, ansl, ansr, "\n")
		}

	}
}
