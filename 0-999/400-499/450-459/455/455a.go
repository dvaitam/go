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
	a := make([]int64, n)
	m := map[int64]int64{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]]++
	}
	ans := make([]int64, 100001)
	for i := int64(1); i <= 100000; i++ {
		if i == 1 {
			ans[i] = m[1]
		} else if i == 2 {
			ans[i] = m[2] * 2
		} else {
			ans[i] = max(ans[i-2], ans[i-3]) + m[i]*i
		}
	}
	total := ans[0]
	for i := 1; i <= 100000; i++ {
		total = max(total, ans[i])
	}
	write(f, total, "\n")

}
