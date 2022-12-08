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
		var x int64
		fmt.Fscan(reader, &n, &x)
		a := make([]int64, n)
		m := map[int64]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]]++
		}
		//fmt.Println(m)
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		ans := 0
		for i := 0; i < n; i++ {
			if m[a[i]] > 0 {
				if m[a[i]*x] > 0 {
					m[a[i]*x]--
				} else {
					ans++
				}
				m[a[i]]--
			}
		}
		write(f, ans, "\n")
	}
}
