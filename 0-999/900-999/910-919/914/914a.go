// 00
package main

import (
	"bufio"
	"fmt"
	"math"
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
	ans := -1000000000
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		sq := int(math.Sqrt(float64(a[i])))
		if sq*sq != a[i] && (sq-1)*(sq-1) != a[i] && (sq+1)*(sq+1) != a[i] {
			ans = max(ans, a[i])
		}
	}
	write(f, ans, "\n")
}
