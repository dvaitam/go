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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s string
	fmt.Fscan(reader, &s)
	n := len(s)
	counts := make([][]int, n)
	for i := 0; i < n; i++ {
		counts[i] = make([]int, 26)
		if i > 0 {
			for j := 0; j < 26; j++ {
				counts[i][j] = counts[i-1][j]
			}
		}
		counts[i][s[i]-'a']++
	}

	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		l--
		r--

		if l == r {
			write(f, "Yes\n")
		} else {
			m := map[int]int{}
			mi := 100000000
			for j := 0; j < 26; j++ {
				res := counts[r][j]
				if l > 0 {
					res -= counts[l-1][j]
				}
				if res > 0 {
					m[j] = res
					mi = min(mi, res)
				}
			}
			if len(m) > 2 {
				write(f, "Yes\n")
			} else if len(m) == 2 {
				if s[l] == s[r] {
					write(f, "No\n")
				} else {
					write(f, "Yes\n")
				}
			} else {
				write(f, "No\n")
			}

		}

	}
}
