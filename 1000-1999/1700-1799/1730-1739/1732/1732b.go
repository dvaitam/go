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
		var s string
		fmt.Fscan(reader, &n, &s)
		red := make([]byte, 0)
		for i := 0; i < n; i++ {
			if len(red) == 0 || s[i] != red[len(red)-1] {
				red = append(red, s[i])
			}
		}
		if len(red) == 1 {
			write(f, "0\n")
		} else {
			if red[0] == '0' {
				write(f, len(red)-2, "\n")
			} else {
				write(f, len(red)-1, "\n")
			}
		}
	}
}
