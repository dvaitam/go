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
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]]++
	}
	ans := 1
	for k := range m {
		tmp := 0
		for j := k; j <= k+5; j++ {
			tmp += m[j]
		}
		ans = max(ans, tmp)
	}
	write(f, ans, "\n")

}
