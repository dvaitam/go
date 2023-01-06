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
	var n, T int
	var c float64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &T, &c)
	a := make([]int64, n)
	cs := make([]int64, n)
	mean := make([]float64, n)
	curr_mean := 0.0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if i == 0 {
			cs[i] = a[i]
		} else {
			cs[i] = cs[i-1] + a[i]
		}
		mean[i] = (curr_mean + float64(a[i])/float64(T)) / c
		curr_mean = mean[i]
	}
	var m int
	fmt.Fscan(reader, &m)
	p := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
		p[i]--
		sm := cs[p[i]]
		if p[i]+1-T > 0 {
			sm -= cs[p[i]-T]
		}
		real := float64(sm) / float64(T)
		write(f, real, mean[p[i]], abs(real-mean[p[i]])/real, "\n")
	}

}
