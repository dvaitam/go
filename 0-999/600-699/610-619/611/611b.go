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
	var a, b int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &a, &b)
	y := make([]int64, 0)
	for bits := 2; bits <= 62; bits++ {
		base := int64((1 << bits) - 1)
		for bit := 0; bit < bits-1; bit++ {
			y = append(y, int64(base^(1<<bit)))
		}
	}
	ans := 0
	for i := 0; i < len(y); i++ {
		if y[i] >= a && y[i] <= b {
			ans++
		}
	}
	write(f, ans, "\n")
}
