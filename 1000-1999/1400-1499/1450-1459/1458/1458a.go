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
func gcd(a int64, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int64, n)
	b := make([]int64, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	gd := int64(0)

	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			if gd == 0 {
				gd = a[i] - a[i-1]
			} else {
				gd = gcd(gd, a[i]-a[i-1])
			}
		}
	}
	for i := 0; i < m; i++ {
		if gd == 0 {
			write(f, b[i]+a[0])
		} else {
			write(f, gcd(gd, b[i]+a[0]))
		}
		if i == m-1 {
			write(f, "\n")
		} else {
			write(f, " ")
		}
	}

}
