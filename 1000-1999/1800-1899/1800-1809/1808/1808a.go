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
func lucky(n int) int {
	mx := 0
	mi := 9
	for n > 0 {
		mx = max(n%10, mx)
		mi = min(n%10, mi)
		n = n / 10
	}
	return mx - mi
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		if r-l > 100 {
			if l%100 > 90 {
				write(f, l-l%100+109, "\n")
			} else {
				write(f, l-l%100+90, "\n")
			}
		} else {
			ll := 0
			ans := -1
			for i := l; i <= r; i++ {
				l := lucky(i)
				if ans == -1 {
					ans = i
					ll = l
				} else {
					if l > ll {
						ans = i
						ll = l
					}
				}
			}
			write(f, ans, "\n")

		}
	}
}
