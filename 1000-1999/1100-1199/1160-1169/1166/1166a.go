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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, 26)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)
		a[s[0]-'a']++
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if a[i] > 1 {
			one := a[i] / 2
			other := a[i] - one
			if one > 1 {
				ans += (one * (one - 1)) / 2
			}
			if other > 1 {
				ans += (other * (other - 1)) / 2
			}
			//write(f, ans, "\n")
		}
	}
	//write(f, a, "\n")
	write(f, ans, "\n")
}
