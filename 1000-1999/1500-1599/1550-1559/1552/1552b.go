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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([][]int, n)
		b := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, 5)
			b[i] = make([]int, 5)
			for j := 0; j < 5; j++ {
				fmt.Fscan(reader, &a[i][j])
				b[i][j] = a[i][j]
			}
		}
		sort.Slice(a, func(i, j int) bool {
			count := 0
			for k := 0; k < 5; k++ {
				if a[i][k] < a[j][k] {
					count++
				}
			}
			return count >= 3
		})
		if n == 1 {
			write(f, "1\n")
		} else {
			min_count := 5
			for i := 1; i < n; i++ {
				count := 0
				for k := 0; k < 5; k++ {
					if a[0][k] < a[i][k] {
						count++
					}
				}
				min_count = min(min_count, count)
			}

			if min_count >= 3 {
				for i := 0; i < n; i++ {
					ok := true
					for j := 0; j < 5; j++ {
						if a[0][j] != b[i][j] {
							ok = false
							break
						}
					}
					if ok {
						write(f, i+1, "\n")
						break
					}
				}
			} else {
				write(f, "-1\n")
			}
		}
	}
}
