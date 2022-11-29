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
func bisect_right(a []int, x int) int {
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if x < a[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
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
		sm := 0
		m := map[int][]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			sm += a[i]
			if _, ok := m[a[i]]; ok {
				m[a[i]] = append(m[a[i]], i)
			} else {
				m[a[i]] = make([]int, 0)
				m[a[i]] = append(m[a[i]], i)
			}
		}
		if (2*sm)%n != 0 {
			write(f, "0\n")
		} else {
			target := (2 * sm) / n
			ans := 0
			for i := 0; i < n; i++ {
				res := m[target-a[i]]
				if len(res) > 0 {
					ans += len(res) - bisect_right(res, i)
				}
			}
			write(f, ans, "\n")
		}

	}
}
