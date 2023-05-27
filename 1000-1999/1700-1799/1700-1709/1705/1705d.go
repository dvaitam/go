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
		var n int
		var s, t string
		fmt.Fscan(reader, &n, &s, &t)
		if s[0] != t[0] || s[n-1] != t[n-1] {
			write(f, "-1\n")
			continue
		}
		ss := make([]int, 0)
		for i := 0; i < n; i++ {
			if i == 0 {
				ss = append(ss, 1)
			} else {
				if s[i] == s[i-1] {
					ss[len(ss)-1]++
				} else {
					ss = append(ss, 1)
				}
			}
		}
		ts := make([]int, 0)
		for i := 0; i < n; i++ {
			if i == 0 {
				ts = append(ts, 1)
			} else {
				if t[i] == t[i-1] {
					ts[len(ts)-1]++
				} else {
					ts = append(ts, 1)
				}
			}
		}
		if len(ss) != len(ts) {
			write(f, "-1\n")
			continue
		}
		diff := make([]int, len(ss))
		for i := 0; i < len(ss); i++ {
			diff[i] = ss[i] - ts[i]
		}
		ans := 0
		for i := 0; i < len(diff); i++ {
			if diff[i] != 0 {
				ans += abs(diff[i])
				diff[i+1] += diff[i]
				diff[i] = 0
			}
		}
		write(f, ans, "\n")
	}
}
