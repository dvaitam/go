// 00
package main

import (
	"bufio"
	"fmt"
	"math"
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
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
		var n, x1, y1 int
		fmt.Fscan(reader, &n)
		x := make([]int, 0)
		y := make([]int, 0)
		for i := 0; i < n*2; i++ {
			fmt.Fscan(reader, &x1, &y1)
			if x1 == 0 {
				y = append(y, y1)
			} else {
				x = append(x, x1)
			}
		}
		sort.Slice(x, func(i, j int) bool { return abs(x[i]) < abs(x[j]) })
		sort.Slice(y, func(i, j int) bool { return abs(y[i]) < abs(y[j]) })
		ans := float64(0)
		for i := 0; i < n; i++ {
			ans += math.Sqrt(float64(x[i]*x[i] + y[i]*y[i]))
		}
		write(f, ans, "\n")

	}
}
