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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		x := make([]int, k)
		c := make([]int, k)
		for i := 0; i < k; i++ {
			fmt.Fscan(reader, &x[i])
		}
		for i := 0; i < k; i++ {
			fmt.Fscan(reader, &c[i])
		}
		s := make([]byte, n)
		chars := make([]byte, 3)
		chars[0] = 'a'
		chars[1] = 'b'
		chars[2] = 'c'
		possible := true
		curr := 0
		for i := 0; i < k; i++ {
			if i == 0 {
				if c[i] > x[i] {
					possible = false
					break
				}
				for j := 0; j < c[i]-3; j++ {
					s[j] = 'd' + byte(i)
				}
				for j := c[i] - 3; j < c[i]; j++ {
					s[j] = chars[curr]
					curr++
					curr %= 3
				}
				for j := c[i]; j < x[i]; j++ {
					s[j] = chars[curr]
					curr++
					curr %= 3
				}
			} else {
				if c[i]-c[i-1] > x[i]-x[i-1] {
					possible = false
					break
				}
				for j := x[i-1]; j < x[i-1]+c[i]-c[i-1]; j++ {
					s[j] = 'd' + byte(i)
				}
				for j := x[i-1] + c[i] - c[i-1]; j < x[i]; j++ {
					s[j] = chars[curr]
					curr++
					curr %= 3
				}

			}
		}
		if possible {
			write(f, "YES\n")
			write(f, string(s), "\n")
		} else {
			write(f, "NO\n")
		}
	}
}
