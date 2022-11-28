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
		var s string
		fmt.Fscan(reader, &n, &s)
		all_strings := map[string]bool{}
		for i := 0; i < n; i++ {
			all_strings[s[i:i+1]] = true
		}
		for i := 0; i < n-1; i++ {
			all_strings[s[i:i+2]] = true
		}
		for i := 0; i < n-2; i++ {
			all_strings[s[i:i+3]] = true
		}
		ok := false
		ans := "a"
		for i := 0; i < 26; i++ {
			st := make([]rune, 1)
			st[0] = rune('a' + i)
			if !all_strings[string(st)] {
				ok = true
				ans = string(st)
				break
			}
		}
		if !ok {
			for i := 0; i < 26; i++ {
				if ok {
					break
				}
				for j := 0; j < 26; j++ {
					st := make([]rune, 2)
					st[0] = rune('a' + i)
					st[1] = rune('a' + j)

					if !all_strings[string(st)] {
						ok = true
						ans = string(st)
						break
					}
				}
			}
		}
		if !ok {
			for i := 0; i < 26; i++ {
				if ok {
					break
				}
				for j := 0; j < 26; j++ {
					if ok {
						break
					}
					for k := 0; k < 26; k++ {
						st := make([]rune, 3)
						st[0] = rune('a' + i)
						st[1] = rune('a' + j)
						st[2] = rune('a' + k)

						if !all_strings[string(st)] {
							ok = true
							ans = string(st)
							break
						}
					}

				}

			}
		}
		write(f, ans, "\n")
	}
}
