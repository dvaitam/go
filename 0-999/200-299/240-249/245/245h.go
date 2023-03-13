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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	n := len(s)
	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]int, n)
		j, k := i, i
		for j >= 0 && k < n {
			if s[j] == s[k] {
				p[j][k]++
				j--
				k++
			} else {
				break
			}
		}
		j, k = i, i+1
		for j >= 0 && k < n {
			if s[j] == s[k] {
				p[j][k]++
				j--
				k++
			} else {
				break
			}
		}
	}
	//write(f, p, "\n")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 && i > 0 {
				p[i][j] += p[i-1][j] + p[i][j-1] - p[i-1][j-1]
			} else if j > 0 {
				p[i][j] += p[i][j-1]

			} else if i > 0 {
				p[i][j] += p[i-1][j]
			}
		}
	}
	//write(f, p, "\n")
	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		l--
		r--
		ans := p[r][r]
		if l > 0 {
			ans -= p[r][l-1]
			ans -= p[l-1][r]
			ans += p[l-1][l-1]
		}

		write(f, ans, "\n")
	}
}
