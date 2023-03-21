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
	m := map[byte]bool{}
	st := "codeforces"
	for i := 0; i < len(st); i++ {
		m[st[i]] = true
	}
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		if m[s[0]] {
			write(f, "Yes\n")
		} else {
			write(f, "No\n")
		}
	}
}
