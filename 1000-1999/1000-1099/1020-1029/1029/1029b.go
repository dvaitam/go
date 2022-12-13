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
	//var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	//fmt.Fscan(reader, &T)
	//for t := 1; t <= T; t++ {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	count := 1
	ans := 1
	for i := 1; i < n; i++ {
		if a[i] <= 2*a[i-1] {
			count++
		} else {
			ans = max(ans, count)
			count = 1
		}
	}
	ans = max(ans, count)
	write(f, ans, "\n")
	//}
}
