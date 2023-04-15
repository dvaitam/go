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
		var s string
		fmt.Fscan(reader, &s)
		ans := 0
		for i := 0; i < len(s); i++ {
			if i > 0 && s[i-1] == '_' && s[i] == '_' {
				ans++
			}
		}
		if s[0] != '^' {
			ans++
		}
		if s[len(s)-1] != '^' {
			ans++
		}
		if len(s) == 1 && s[0] == '^' {
			ans++
		}
		write(f, ans, "\n")
	}
}
