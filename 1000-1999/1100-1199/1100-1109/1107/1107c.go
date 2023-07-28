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
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var s string
	fmt.Fscan(reader, &s)
	b := make([]int64, 0)
	b = append(b, a[0])
	ans := int64(0)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			b = append(b, a[i])
		} else {
			sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
			for j := 0; j < len(b) && j < k; j++ {
				ans += b[j]
			}
			b = make([]int64, 0)
			b = append(b, a[i])
		}
	}
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	//write(f, b, "\n")
	for j := 0; j < len(b) && j < k; j++ {
		ans += b[j]
	}
	write(f, ans, "\n")
}
