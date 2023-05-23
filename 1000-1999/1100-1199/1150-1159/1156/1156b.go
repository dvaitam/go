// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		m := map[byte]int{}
		for i := 0; i < len(s); i++ {
			m[s[i]]++
		}
		first := make([]byte, 0)
		second := make([]byte, 0)
		chars := make([]byte, 0)
		for ch := range m {
			chars = append(chars, ch)
		}
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		for i := 0; i < len(chars); i++ {
			if len(first) == 0 {
				first = append(first, chars[i])
			} else {
				if chars[i] != first[len(first)-1]+1 {
					first = append(first, chars[i])
				} else {
					second = append(second, chars[i])
				}
			}
		}
		if len(second) != 0 && (second[0] == first[len(first)-1]+1 || second[0]+1 == first[len(first)-1]) && (first[0] == second[len(second)-1]+1 || first[0]+1 == second[len(second)-1]) {
			write(f, "No answer\n")
		} else {
			if len(second) != 0 && (second[0] == first[len(first)-1]+1 || second[0]+1 == first[len(first)-1]) {
				for _, ch := range second {
					for i := 0; i < m[ch]; i++ {
						write(f, string(ch))
					}
				}
				for _, ch := range first {
					for i := 0; i < m[ch]; i++ {
						write(f, string(ch))
					}
				}

				write(f, "\n")

			} else {
				for _, ch := range first {
					for i := 0; i < m[ch]; i++ {
						write(f, string(ch))
					}
				}
				for _, ch := range second {
					for i := 0; i < m[ch]; i++ {
						write(f, string(ch))
					}
				}
				write(f, "\n")
			}

		}
	}
}
