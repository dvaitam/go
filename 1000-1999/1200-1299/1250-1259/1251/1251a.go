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
		curr := 1
		m := make([]bool, 26)
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] {
				if curr%2 == 1 {
					m[s[i-1]-'a'] = true
				}
				curr = 1
			} else {
				curr++
			}
		}
		if curr%2 == 1 {
			m[s[len(s)-1]-'a'] = true
		}
		for i := 0; i < 26; i++ {
			if m[i] {
				write(f, string('a'+i))
			}
		}
		write(f, "\n")
	}
}
