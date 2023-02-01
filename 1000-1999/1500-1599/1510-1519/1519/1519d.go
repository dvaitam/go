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
	a := make([]int64, n)
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	total := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
		total += a[i] * b[i]
	}
	ans := total
	for i := 0; i < n; i++ {
		l, r := i, i
		start_total := total
		for l >= 0 && r < n {
			start_total += (a[r] - a[l]) * (b[l] - b[r])
			ans = max(ans, start_total)
			l--
			r++
		}
		if i+1 < n {
			l, r = i, i+1
			start_total = total
			for l >= 0 && r < n {
				start_total += (a[r] - a[l]) * (b[l] - b[r])
				ans = max(ans, start_total)
				l--
				r++
			}
		}
	}
	write(f, ans, "\n")

}
