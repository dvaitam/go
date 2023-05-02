// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var n, d int
	var m int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &d, &m)
	l := make([]int64, 0)
	h := make([]int64, 0)
	for i := 0; i < n; i++ {
		var tmp int64
		fmt.Fscan(reader, &tmp)
		if tmp > m {
			h = append(h, tmp)
		} else {
			l = append(l, tmp)
		}
	}
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	sort.Slice(h, func(i, j int) bool { return h[i] < h[j] })
	ldp := make([]int64, len(l))
	for i := 0; i < len(l); i++ {
		if i == 0 {
			ldp[i] = l[i]
		} else {
			ldp[i] = ldp[i-1] + l[i]
		}
	}
	hdp := make([]int64, len(h))
	for i := len(h) - 1; i >= 0; i-- {
		if i == len(h)-1 {
			hdp[i] = h[i]
		} else {
			hdp[i] = hdp[i+1] + h[i]
		}
	}
	dp := make([]int64, len(l)+1)
	for i := 0; i <= len(l); i++ {
		tmp := int64(0)
		if len(l) > 0 {
			tmp += ldp[len(l)-1]
		}
		if i > 0 {
			tmp -= ldp[i-1]
		}
		rem := i + len(h)
		total := rem / (d + 1)
		if rem%(d+1) != 0 {
			total++
		}
		if len(hdp)-total >= 0 && len(hdp)-total < len(hdp) {
			tmp += hdp[len(hdp)-total]
		}
		dp[i] = tmp
	}
	ans := dp[0]
	for i := 1; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}

	write(f, ans, "\n")

}
