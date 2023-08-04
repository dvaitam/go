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
func gcd(a int64, b int64) int64 {
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
	x := make([]int64, n)
	p := make([]int64, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
	}
	ans := x[1] - x[0]
	for i := 2; i < n; i++ {
		ans = gcd(ans, x[i]-x[i-1])
	}
	ans_i := -1
	for i := 0; i < m; i++ {
		if ans%p[i] == 0 {
			ans_i = i + 1
			break
		}
	}
	if ans_i == -1 {
		write(f, "NO\n")
	} else {
		write(f, "YES\n")
		write(f, x[0], ans_i, "\n")
	}

}
