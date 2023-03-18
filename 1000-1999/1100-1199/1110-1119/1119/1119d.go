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
func bisect_left(a []int64, x int64) int {
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
func main() {
	var n, q int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	s = append(s, s[n-1]+1000000000000000001)
	d := make([]int64, 0)
	for i := 1; i <= n; i++ {
		d = append(d, s[i]-s[i-1])
	}
	sort.Slice(d, func(i, j int) bool { return d[i] < d[j] })

	ds := make([]int64, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ds[i] = d[i]
		} else {
			ds[i] = ds[i-1] + d[i]
		}
	}
	//write(f, d, "\n")
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var l, r int64
		fmt.Fscan(reader, &l, &r)
		w := r - l + 1
		//write(f, w, "\n")
		j := bisect_left(d, w)
		if j > 0 {
			write(f, ds[j-1]+int64(n-j)*w, " ")
		} else {
			write(f, int64(n)*w, " ")
		}
	}
	write(f, "\n")

}
