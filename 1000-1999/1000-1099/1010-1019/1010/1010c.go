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
	var n, k int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int64, n)
	ans := int64(-1)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i] = a[i] % k
		if a[i] != 0 {
			if ans == -1 {
				ans = a[i]
			} else {
				ans = gcd(ans, a[i])
			}
		}
	}
	if ans == -1 {
		ans = 0
	}
	ans = gcd(ans, k)
	if ans == 0 {
		write(f, "1\n0\n")
	} else {
		write(f, k/ans, "\n")
		write(f, "0")
		for i := int64(1); i < k/ans; i++ {
			write(f, " ", i*ans)
		}
		write(f, "\n")
	}

}
