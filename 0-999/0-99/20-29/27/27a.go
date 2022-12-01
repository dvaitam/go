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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	m := map[int]bool{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]] = true
	}
	ans := -1
	for i := 1; i <= 3001; i++ {
		if !m[i] {
			ans = i
			break
		}
	}
	write(f, ans, "\n")

}
