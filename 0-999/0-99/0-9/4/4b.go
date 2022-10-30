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
	var n, s int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &s)
	mi := make([]int, n)
	mx := make([]int, n)
	mi_sum := 0
	mx_sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &mi[i], &mx[i])
		mi_sum += mi[i]
		mx_sum += mx[i]
	}
	if s >= mi_sum && s <= mx_sum {
		write(f, "YES\n")
		s -= mi_sum
		for i := 0; i < n; i++ {
			if mx[i]-mi[i] <= s {
				s -= (mx[i] - mi[i])
				mi[i] = mx[i]
			} else if s > 0 {
				mi[i] += s
				s = 0
			}
			write(f, mi[i])
			if i == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}

		}
	} else {
		write(f, "NO\n")
	}
}
