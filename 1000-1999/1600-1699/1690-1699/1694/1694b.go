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
		var s string
		fmt.Fscan(reader, &n, &s)
		ans := int64(0)
		for i := 1; i < int(n); i++ {
			if s[i] == '1' && s[i-1] == '1' {
				ans += int64(i)
			}
			if s[i] == '0' && s[i-1] == '0' {
				ans += int64(i)
			}
		}

		ans = (n*(n+1))/2 - ans
		write(f, ans, "\n")
	}
}
