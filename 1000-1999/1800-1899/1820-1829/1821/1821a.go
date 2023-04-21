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
		ans := 1
		for i := 0; i < len(s); i++ {
			if i == 0 {
				if s[i] == '?' {
					ans *= 9
				} else if s[i] == '0' {
					ans *= 0
				}
			} else {
				if s[i] == '?' {
					ans *= 10
				}
			}
		}
		write(f, ans, "\n")
	}
}
