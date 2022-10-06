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

type Pair struct {
	i, v int
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
		a := make([]Pair, n)
		as := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = Pair{i: i}
			fmt.Fscan(reader, &a[i].v)
			as[i] = a[i].v
		}
		sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })
		x := make([]int, n+1)
		curr := 1
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				x[a[i].i+1] = curr
			} else {
				x[a[i].i+1] = -curr
			}
			if i%2 == 1 {
				curr++
			}
		}
		ans := int64(0)
		for i := 1; i <= n; i++ {
			ans += abs(int64(as[i-1] * 2 * x[i]))
		}
		write(f, ans, "\n")
		for i := 0; i <= n; i++ {
			write(f, x[i])
			if i == n {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	}
}
