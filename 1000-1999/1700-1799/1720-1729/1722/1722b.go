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
		fmt.Fscan(reader, &n)
		var s1, s2 string
		fmt.Fscan(reader, &s1, &s2)
		ok := true
		for i := 0; i < n; i++ {
			if s1[i] == 'R' && s2[i] != 'R' || s2[i] == 'R' && s1[i] != 'R' {
				ok = false
				break
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
