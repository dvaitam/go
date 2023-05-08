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
		var s string
		fmt.Fscan(reader, &s)
		start := make([]bool, len(s))
		front := make([]bool, len(s))
		back := make([]bool, len(s))
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				front[i] = true
			}
		}
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				back[i] = true
			}
		}
		l := 2
		for len(start) > 1 {
			if l%2 == 0 {
				start = make([]bool, len(s)-l+1)
				for i := 0; i < len(s)-l+1; i++ {
					start[i] = front[i] && back[i+1]
				}
			} else {
				front = make([]bool, len(s))
				back = make([]bool, len(s))
				for i := 0; i < len(s)-l; i++ {
					if s[i] == s[i+l] && s[i+l] == s[i+l-1] {
						front[i] = start[i] || start[i+1]
					} else if s[i] == s[i+l] {
						front[i] = start[i+1]
					} else if s[i+l] == s[i+l-1] {
						front[i] = start[i]
					} else {
						front[i] = false
					}
				}
				for i := 1; i < len(s)-l+1; i++ {
					if s[i] == s[i-1] && s[i-1] == s[i+l-1] {
						back[i] = start[i+1] || start[i]
					} else if s[i] == s[i-1] {
						back[i] = start[i+1]
					} else if s[i-1] == s[i+l-1] {
						back[i] = start[i]
					} else {
						back[i] = false
					}
				}
			}
			l++
		}

		if start[0] {
			write(f, "Draw\n")
		} else {
			write(f, "Alice\n")
		}
	}
}
