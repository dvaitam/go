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
func get_string(ll int, prev string, match int) string {
	s := make([]byte, ll)
	for i := 0; i < ll; i++ {
		if i < match {
			s[i] = prev[i]
		} else {
			if prev[i] == 'a' {
				s[i] = 'b'
			} else {
				s[i] = 'a'
			}
		}

	}
	return string(s)

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
		a := make([]int, n)
		mx := 1
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			mx = max(mx, a[i])
		}
		st := make([]byte, mx)
		for i := 0; i < mx; i++ {
			st[i] = 'a'
		}
		start := string(st)
		prev := start
		for i := 0; i < n+1; i++ {
			curr := start
			if i > 0 {
				curr = get_string(mx, prev, a[i-1])
			}
			write(f, curr, "\n")
			prev = curr
		}
	}
}
