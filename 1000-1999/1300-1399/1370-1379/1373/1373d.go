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
func get_max_subarray(a []int64) int64 {
	n := len(a)
	if n == 0 {
		return 0
	}
	dpi := make([]int64, n)
	dpe := make([]int64, n)
	dpi[0] = a[0]

	for i := 1; i < n; i++ {
		dpi[i] = max(a[i], dpi[i-1]+a[i])
		dpe[i] = max(dpi[i-1], dpe[i-1])
	}
	return max(dpe[n-1], dpi[n-1])
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		total := int64(0)
		a1 := make([]int64, 0)
		a2 := make([]int64, 0)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i%2 == 0 {
				total += a[i]
			}
			if i > 0 {
				if i%2 == 0 {
					a2 = append(a2, a[i-1]-a[i])
				} else {
					a1 = append(a1, a[i]-a[i-1])
				}
			}
		}
		ans := total + max(get_max_subarray(a1), get_max_subarray(a2))
		write(f, ans, "\n")
	}
}
