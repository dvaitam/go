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
func find_min(a []int) int {
	mi := 1000000000
	for i := 0; i < len(a); i++ {
		mi = min(mi, a[i])
	}
	return mi
}
func main() {
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &k)
	h := make([][]int, n)
	for i := 0; i < n; i++ {
		h[i] = make([]int, m-1)
		for j := 0; j < m-1; j++ {
			fmt.Fscan(reader, &h[i][j])
		}
	}
	v := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		v[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &v[i][j])
		}
	}
	if k%2 == 1 {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				write(f, "-1 ")
			}
			write(f, "\n")
		}
	} else {
		c := make([][][]int, n)
		for i := 0; i < n; i++ {
			c[i] = make([][]int, m)
			for j := 0; j < m; j++ {
				c[i][j] = make([]int, k/2)
			}
		}
		for p := 0; p < k/2; p++ {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if p == 0 {
						a := make([]int, 0)
						if i > 0 {
							a = append(a, v[i-1][j])
						}
						if i < n-1 {
							a = append(a, v[i][j])
						}
						if j > 0 {
							a = append(a, h[i][j-1])
						}
						if j < m-1 {
							a = append(a, h[i][j])
						}
						c[i][j][p] = find_min(a)
					} else {
						a := make([]int, 0)
						if i > 0 {
							a = append(a, c[i-1][j][p-1]+v[i-1][j])
						}
						if i < n-1 {
							a = append(a, c[i+1][j][p-1]+v[i][j])
						}
						if j > 0 {
							a = append(a, c[i][j-1][p-1]+h[i][j-1])
						}
						if j < m-1 {
							a = append(a, c[i][j+1][p-1]+h[i][j])
						}
						c[i][j][p] = find_min(a)
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				write(f, 2*c[i][j][(k/2)-1], " ")
			}
			write(f, "\n")
		}
	}

}
