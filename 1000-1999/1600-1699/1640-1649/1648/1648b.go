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
		var n, c int
		fmt.Fscan(reader, &n, &c)
		cnt := make([]int, c+1)
		sum := make([]int, c+1)
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(reader, &x)
			cnt[x]++
		}
		for i := 1; i <= c; i++ {
			sum[i] = sum[i-1] + cnt[i]
		}
		failed := false
		for y := 1; y <= c; y++ {
			if cnt[y] == 0 {
				continue
			}
			for i := 1; i*y <= c; i++ {
				r := min(c, i*y+y-1)
				if sum[r]-sum[i*y-1] > 0 {
					if cnt[i] == 0 {
						write(f, "No\n")
						failed = true
						break
					}
				}
			}
			if failed {
				break
			}
		}
		if !failed {
			write(f, "Yes\n")
		}
	}
}
