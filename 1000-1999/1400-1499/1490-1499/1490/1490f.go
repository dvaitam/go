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
		m := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]]++
		}
		uniq := map[int]bool{}
		for _, v := range m {
			uniq[v] = true
		}
		ans := n
		for k := range uniq {
			rem := 0
			for _, val := range m {
				if val < k {
					rem += val
				} else if val > k {
					rem += val - k
				}
			}
			ans = min(ans, rem)
		}
		write(f, ans, "\n")
	}
}
