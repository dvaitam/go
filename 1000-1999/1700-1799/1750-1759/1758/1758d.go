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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int64
		fmt.Fscan(reader, &n)
		ans := make([]int64, 0)
		ans = append(ans, 3*n)
		ans = append(ans, 5*n)
		rem := 16*n*n - 8*n
		left, right := int64(0), int64(0)
		if n%2 == 1 {
			ans = append(ans, 4*n)
			rem -= 4 * n
			left, right = 4*n-1, 4*n+1
		} else {
			left, right = 4*n-1, 4*n+1
		}
		for rem > 0 {
			ans = append(ans, left)
			ans = append(ans, right)
			rem -= (left + right)
			left--
			right++
		}
		for i := int64(0); i < n; i++ {
			write(f, ans[i])
			if i == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}

	}
}
