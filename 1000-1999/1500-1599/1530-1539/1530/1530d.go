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
		counts := map[int]int{}
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
				if counts[a[i]] > 1 {
					g = append(g, i+1)
					mg[i+1] = true
				}

			}
		}
		//write(f, g, counts, "\n")
		curr := 1
		ans := make([]int, n)
		if len(g) != 1 {
			for i := 0; i < n; i++ {
				if counts[a[i]] > 1 {
					if mg[i+1] {
						ans[i] = g[curr%len(g)]
						curr++
					}
				}
			}
			for i := 0; i < n; i++ {
				if counts[a[i]] == 1 {
					ans[i] = a[i]
				}
			}
			gifted := map[int]bool{}
			for i := 0; i < n; i++ {
				if ans[i] != 0 {
					gifted[ans[i]] = true
				}
			}
			not_gifted := make([]int, 0)
			for i := 0; i < n; i++ {
				if !gifted[i+1] {
					not_gifted = append(not_gifted, i+1)
				}
			}
			curr := 0
			for i := 0; i < n; i++ {
				if ans[i] == 0 {
					ans[i] = not_gifted[curr]
					curr++
				}
			}

			write(f, len(counts), "\n")
		} else {
			for i := 0; i < n; i++ {
				if counts[a[i]] > 1 && g[0] != i+1 {
					ans[i] = g[0]
				} else {
					ans[i] = a[i]
				}
			}
			write(f, len(counts), "\n")
		}

		for i := 0; i < n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
	}
}
