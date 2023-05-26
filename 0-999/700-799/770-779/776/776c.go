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
func pow(x int64, n int) int64 {
	if n == 0 {
		return 1
	} else {
		return x * pow(x, n-1)
	}
}
func main() {
	var n, k int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int64, n)
	c := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if i == 0 {
			c[i] = a[i]
		} else {
			c[i] = c[i-1] + a[i]
		}
	}
	m := map[int64]int64{}
	m[0] = 1
	mx := pow(10, 15)
	ans := int64(0)
	for i := int64(0); i < n; i++ {
		for j := 0; j < 60; j++ {
			val := pow(k, j)
			if k == 1 && j == 1 || k == -1 && j == 2 || abs(val) > mx {
				break
			}
			ans += m[c[i]-val]
		}
		m[c[i]]++
	}
	write(f, ans, "\n")

}
