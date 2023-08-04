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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]]++
	}
	var mn int
	fmt.Fscan(reader, &mn)
	b := make([]int, mn)
	c := make([]int, mn)
	for i := 0; i < mn; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < mn; i++ {
		fmt.Fscan(reader, &c[i])
	}
	ans := 0
	for i := 0; i < mn; i++ {
		if m[b[i]] > m[b[ans]] {
			ans = i
		} else if m[b[i]] == m[b[ans]] {
			if m[c[i]] > m[c[ans]] {
				ans = i
			}
		}
	}
	write(f, ans+1, "\n")

}
