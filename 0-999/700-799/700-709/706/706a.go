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
	var a, b, n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &a, &b, &n)
	ans := float64(-1)
	for i := 0; i < n; i++ {
		var x, y, v int
		fmt.Fscan(reader, &x, &y, &v)
		if ans == -1 {
			ans = math.Sqrt(float64((x-a)*(x-a)+(y-b)*(y-b))) / float64(v)
		} else {
			ans = min(ans, math.Sqrt(float64((x-a)*(x-a)+(y-b)*(y-b)))/float64(v))
		}
	}
	write(f, ans, "\n")
}
