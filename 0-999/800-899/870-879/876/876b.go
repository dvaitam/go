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
	var n, k, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k, &m)
	a := make([]int, n)
	c := map[int][]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		c[a[i]%m] = append(c[a[i]%m], i)
	}
	ok := false
	for _, val := range c {
		if len(val) >= k {
			ok = true
			write(f, "Yes\n")
			for i := 0; i < k; i++ {
				write(f, a[val[i]], " ")
			}
			write(f, "\n")

			break
		}
	}
	if !ok {
		write(f, "No\n")
	}
}
