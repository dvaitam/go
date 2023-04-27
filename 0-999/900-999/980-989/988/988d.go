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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	m := map[int]bool{}
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		m[x[i]] = true
	}
	ans := make([]int, 1)
	ans[0] = x[0]
	for i := 0; i < n; i++ {
		if len(ans) == 3 {
			break
		}

		for j := 0; j < 31; j++ {
			curr := 1 << j
			if len(ans) < 3 {
				if m[x[i]+curr] && m[x[i]+2*curr] {
					tmp := make([]int, 0)
					tmp = append(tmp, x[i])
					tmp = append(tmp, x[i]+curr)
					tmp = append(tmp, x[i]+2*curr)
					ans = tmp
				} else {
					if len(ans) < 2 {
						if m[x[i]+curr] {
							tmp := make([]int, 0)
							tmp = append(tmp, x[i])
							tmp = append(tmp, x[i]+curr)
							ans = tmp
						}
					}
				}
			} else {
				break
			}

		}
	}
	write(f, len(ans), "\n")
	for i := 0; i < len(ans); i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
