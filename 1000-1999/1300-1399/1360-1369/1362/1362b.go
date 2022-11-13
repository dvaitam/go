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
		m := map[int]bool{}

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = true
		}
		found := false
		for i := 1; i < 1024; i++ {
			t := map[int]bool{}
			for j := 0; j < n; j++ {
				t[a[j]^i] = true
			}
			possible := true
			for j := 0; j < n; j++ {
				if !t[a[j]] {
					possible = false
					break
				}
			}
			if possible {
				found = true
				write(f, i, "\n")
				break
			}
		}
		if !found {
			write(f, "-1\n")
		}
	}
}
