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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	var s string
	fmt.Fscan(reader, &s)
	sb := []byte(s)
	ans := 0
	for i := 1; i < n; i++ {
		if sb[i] == '.' {
			if sb[i-1] == '.' {
				ans++
			}
		}
	}
	for i := 0; i < m; i++ {
		var j int
		var t string
		fmt.Fscan(reader, &j, &t)
		j--
		if t[0] == '.' {
			if sb[j] != '.' {
				if j > 0 && sb[j-1] == '.' {
					ans++
				}
				if j+1 < n && sb[j+1] == '.' {
					ans++
				}
				sb[j] = t[0]
			}
		} else {
			if sb[j] == '.' {
				if j > 0 && sb[j-1] == '.' {
					ans--
				}
				if j+1 < n && sb[j+1] == '.' {
					ans--
				}
			}
			sb[j] = t[0]
		}
		write(f, ans, "\n")
	}
}
