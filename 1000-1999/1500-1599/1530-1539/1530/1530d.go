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
		counts := make([]int, n+1)
		b := make([]bool, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			counts[a[i]]++
			b[a[i]-1] = true
		}
		g := make([]int, 0)
		mg := map[int]bool{}
		for i := 0; i < n; i++ {
			if !b[i] {
				g = append(g, i+1)
				mg[i+1] = true
			}
		}
		//write(f, g, "\n")
		curr := 0
		bs := make([]bool, n)
		ans := make([]int, n)
		fill := 0
		for i := 0; i < n; i++ {
			if !bs[a[i]-1] {
				if counts[a[i]] == 1 || a[a[i]-1] != i+1 {
					ans[i] = a[i]
					bs[a[i]-1] = true
					fill++
				}

			}
		}
		//write(f, ans, "\n")
		for i := 0; i < n; i++ {
			if ans[i] == 0 {
				if mg[i+1] {
					if g[len(g)-1] == i+1 {
						g[len(g)-1], g[len(g)-2] = g[len(g)-2], g[len(g)-1]
					}
					ans[i] = g[len(g)-1]
					g = g[:len(g)-1]
				}
			}

		}
		for i := 0; i < n; i++ {
			if ans[i] == 0 {
				ans[i] = g[curr]
				curr++
			}
		}
		write(f, fill, "\n")
		for i := 0; i < n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
	}
}
