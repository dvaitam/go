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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		n := len(s)
		ans := n
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				first, second := '0'+byte(i), '0'+byte(j)
				flip := false
				count := 0
				for k := 0; k < n; k++ {
					if flip {
						if second == s[k] {
							count++
							flip = !flip
						}
					} else {
						if first == s[k] {
							count++
							flip = !flip
						}
					}
				}
				if first != second && count%2 == 1 {
					count -= 1
				}
				ans = min(ans, n-count)
			}
		}
		write(f, ans, "\n")
	}
}
