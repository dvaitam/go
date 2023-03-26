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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		b := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i], &b[i])
		}
		curr := b[0] * a[0]
		lm := b[0]
		ans := 1
		for i := 1; i < n; i++ {
			curr = gcd(curr, b[i]*a[i])
			lm = lcm(lm, b[i])
			if curr%lm != 0 {
				//	write(f, curr, "\n")
				ans++
				curr = b[i] * a[i]
				lm = b[i]
			}
		}
		write(f, ans, "\n")
	}
}
