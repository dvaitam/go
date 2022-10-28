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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		b := make([]int, 0)
		b = append(b, a[0])
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				b = append(b, a[i])
			}
		}
		m := map[int]int{}
		m[b[0]] = 0
		if len(b) > 1 {
			m[b[0]] = 1
		}
		for i := 1; i < len(b); i++ {
			if m[b[i]] > 0 || i == len(b)-1 {
				if i == len(b)-1 && (b[i] == b[0] || m[b[i]] > 0) {
					m[b[i]] += 0
				} else {
					m[b[i]] += 1
				}
			} else {
				m[b[i]] = 2
			}
		}
		ans := m[b[0]]
		for _, v := range m {
			ans = min(ans, v)
		}
		write(f, ans, "\n")
	}
}
