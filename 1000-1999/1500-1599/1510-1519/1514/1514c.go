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
	var n int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	ans := make([]int64, 0)
	ans = append(ans, 1)

	for i := int64(2); i < n-1; i++ {
		if gcd(n, i) == 1 {
			ans = append(ans, i)
		}
	}

	rem := int64(1)
	for i := 0; i < len(ans); i++ {
		rem = rem * ans[i]
		rem = rem % n
	}
	//write(f, rem, "\n")

	if rem == n-1 && n-1 != 1 {
		ans = append(ans, n-1)
	}

	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
	write(f, len(ans), "\n")
	for i := 0; i < len(ans); i++ {
		write(f, ans[i])
		if i == len(ans)-1 {
			write(f, "\n")
		} else {
			write(f, " ")
		}
	}
}
