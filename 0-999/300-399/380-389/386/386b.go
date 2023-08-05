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
	var n, T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	sort.Ints(t)
	fmt.Fscan(reader, &T)
	ans := 0
	i := 0
	j := 0
	count := 0
	for j < n {
		for j < n && t[j]-t[i] <= T {
			count++
			j++
		}
		ans = max(ans, count)
		if j == n {
			break
		}
		for i < n && t[j]-t[i] >= T {
			i++
			count--
		}
	}
	write(f, ans, "\n")
}
