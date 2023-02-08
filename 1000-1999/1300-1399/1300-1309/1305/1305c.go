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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func get_val(r int, i int, n int, m int, cache map[int]map[int]int, b []int) int {
	if i == n-1 {
		return (r + b[i]) % m
	}
	if cache[r] != nil {
		if cache[r][i] > 0 {
			return cache[r][i] - 1
		}
	}

	next_val := ((r + b[i]) % m) * get_val((r+b[i])%m, i+1, n, m, cache, b)
	if cache[r] == nil {
		cache[r] = map[int]int{}
	}
	cache[r][i] = next_val + 1
	return next_val
}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans *= abs(a[i] - a[j])
			ans = ans % m
		}
	}
	write(f, ans, "\n")
}
