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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, n)
		}
		start_i := 0
		d1 := 0
		start_j := n - 1
		d2 := n - 1
		top := true
		for k > 0 {
			if top {
				for i := start_i; i < n && k > 0; i++ {
					a[i][i-d1] = 1
					k--
				}
				d1++
				start_i++
			} else {
				for j := start_j; j < n && k > 0; j++ {
					a[j-d2][j] = 1
					k--
				}
				d2--
				start_j--
			}
			top = !top
		}
		min_r := 100000
		min_c := 100000
		max_r := 0
		max_c := 0
		for i := 0; i < n; i++ {
			s := 0
			for j := 0; j < n; j++ {
				s += a[i][j]
			}
			min_r = min(min_r, s)
			max_r = max(max_r, s)
		}
		for i := 0; i < n; i++ {
			s := 0
			for j := 0; j < n; j++ {
				s += a[j][i]
			}
			min_c = min(min_c, s)
			max_c = max(max_c, s)
		}
		write(f, (max_r-min_r)*(max_r-min_r)+(max_c-min_c)*(max_c-min_c), "\n")
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				write(f, a[i][j])
			}
			write(f, "\n")
		}
	}
}
