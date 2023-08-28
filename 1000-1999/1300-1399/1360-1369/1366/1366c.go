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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(reader, &a[i][j])
			}
		}
		ans := 0
		for i := 0; i < m+n; i++ {
			j := m + n - 2 - i
			if i >= j {
				continue
			}
			mp := make([]bool, 100000)
			zeros := 0
			ones := 0
			for k := 0; k <= i; k++ {
				if k >= n || i-k >= m {
					continue
				}
				//write(f, k, i-k, "count\n")
				if !mp[k*100+i-k] {
					mp[k*100+i-k] = true
					if a[k][i-k] == 1 {
						ones++
					} else {
						zeros++
					}
				}
			}
			for k := min(n-1, j); k >= 0; k-- {
				if j-k >= m || k >= n {
					continue
				}
				//write(f, k, j-k, "count\n")

				if !mp[k*100+j-k] {
					mp[k*100+j-k] = true
					if a[k][j-k] == 1 {
						ones++
					} else {
						zeros++
					}
				}
			}
			//	write(f, mp, "\n")
			//write(f, zeros, ones, i, j, "\n")
			ans += min(zeros, ones)
		}
		write(f, ans, "\n")
	}

}
