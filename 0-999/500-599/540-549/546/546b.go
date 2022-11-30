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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	ans := 0
	sa := make([]int, n)
	sa[0] = a[0]
	for i := 1; i < n; i++ {
		if a[i] > sa[i-1] {
			sa[i] = a[i]
		} else {
			diff := sa[i-1] - a[i]
			sa[i] = a[i] + diff + 1
			ans += diff + 1
		}
	}
	// write(f, a, "\n")
	// write(f, sa, "\n")
	write(f, ans, "\n")
}
