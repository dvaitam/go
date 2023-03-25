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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	m := make([]int, 1024)
	for i := 0; i < 1024; i++ {
		m[i] = 1000
	}
	for i := 0; i < n; i++ {
		m[a[i]] = min(m[a[i]], a[i])
		for j := 0; j < 1024; j++ {
			if m[j] < a[i] {
				nxt := j ^ a[i]
				m[nxt] = min(m[nxt], a[i])
			}
		}
	}
	ans := make([]int, 1)
	for i := 1; i < 1024; i++ {
		if m[i] != 1000 {
			ans = append(ans, i)
		}
	}
	//sort.Ints(ans)
	write(f, len(ans), "\n")
	for i := 0; i < len(ans); i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
