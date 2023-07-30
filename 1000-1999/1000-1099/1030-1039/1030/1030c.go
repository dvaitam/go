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
	var s string
	fmt.Fscan(reader, &n, &s)
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		a = append(a, int(s[i]-'0'))
	}
	sm := 0
	for i := 0; i < n; i++ {
		sm += a[i]
	}
	ok := false
	for d := 1; d < sm; d++ {
		if sm%d != 0 {
			continue
		}
		curr := 0
		i := 0
		for i < n {
			for i < n && curr < d {
				curr += a[i]
				i++
			}
			if curr != d {
				break
			}
			curr = 0
			//fmt.Println(i, curr, d)

		}
		if i == n && curr == 0 {
			ok = true
			break
		}
	}
	if ok || sm == 0 {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}

}
