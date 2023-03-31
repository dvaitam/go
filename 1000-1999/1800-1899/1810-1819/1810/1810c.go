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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var c, d int64
		fmt.Fscan(reader, &n, &c, &d)
		a := make([]int, n)
		ans := c*int64(n) + d
		m := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]]++
		}
		keys := make([]int, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		delete := 0
		so_far := 0
		for i := 0; i < len(keys); i++ {
			if m[keys[i]] > 1 {
				delete += m[keys[i]] - 1
			}
			so_far += m[keys[i]]
			ans = min(ans, int64(delete+n-so_far)*c+int64(keys[i]-i-1)*d)
		}
		write(f, ans, "\n")
	}
}
