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
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		b := make([]int, n)
		s := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			s += a[i]
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
			s += b[i]
		}
		ff := make([]bool, s+1)
		ff[0] = true
		for i := 0; i < n; i++ {
			for j := s; j >= 0; j-- {
				if ff[j] {
					if j+a[i] <= s {
						ff[j+a[i]] = true
					}
					if j+b[i] <= s {
						ff[j+b[i]] = true
					}
					ff[j] = false
				}
			}
		}
		best := -1
		for j := 0; j <= s; j++ {
			if ff[j] {
				if abs(2*j-s) < abs(2*best-s) {
					best = j
				}
			}

		}
		sqaure_sum := 0
		for i := 0; i < n; i++ {
			sqaure_sum += a[i]*a[i] + b[i]*b[i]
		}
		ans := best*best + (s-best)*(s-best)
		ans += sqaure_sum * (n - 2)
		write(f, ans, "\n")
	}
}
