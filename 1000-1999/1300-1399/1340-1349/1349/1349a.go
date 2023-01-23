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
func lcm(a int64, b int64) int64 {
	return a * b / gcd(a, b)
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	gdi := make([]int64, n)
	gdi[0] = a[0]
	for i := 1; i < n; i++ {
		gdi[i] = gcd(a[i], gdi[i-1])
	}
	gdr := make([]int64, n)
	gdr[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		gdr[i] = gcd(a[i], gdr[i+1])
	}

	gde := make([]int64, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			gde[i] = gdr[1]
		} else if i == n-1 {
			gde[i] = gdi[n-2]
		} else {
			gde[i] = gcd(gdi[i-1], gdr[i+1])
		}
	}
	ans := gde[0]
	for i := 1; i < n; i++ {
		ans = lcm(ans, gde[i])
	}
	write(f, ans, "\n")
}
