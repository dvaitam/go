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

var found bool
var s1, s2 []byte
var as1, as2 string
var n int

func backtrack(s string, i int, ub1 int, ub2 int) {
	if i == n && ub1 == 0 && ub2 == 0 {
		found = true
		as1 = string(s1)
		as2 = string(s2)
		// fmt.Println("Found")
		return
	}
	if i == n {
		return
	}
	if found {
		return
	}
	if n-i > ub1 {
		s1[i] = '('
		if s[i] == '0' && ub2 > 0 {
			s2[i] = ')'
			if !found {
				backtrack(s, i+1, ub1+1, ub2-1)
			}
		}
		if s[i] == '1' {
			s2[i] = '('
			if !found {
				backtrack(s, i+1, ub1+1, ub2+1)

			}
		}
	}

	if ub1 > 0 {
		s1[i] = ')'
		if s[i] == '0' {
			s2[i] = '('
			if !found {
				backtrack(s, i+1, ub1-1, ub2+1)

			}
		}
		if s[i] == '1' && ub2 > 0 {
			s2[i] = ')'
			if !found {
				backtrack(s, i+1, ub1-1, ub2-1)
			}
		}
	}

}
func balanced(s1 []byte) bool {
	count := 0
	for i := 0; i < n; i++ {
		if s1[i] == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &n, &s)
		if s[0] != '1' || s[n-1] != '1' {
			write(f, "NO\n")
			continue
		}
		eq := make([]int, 0)
		neq := make([]int, 0)
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				eq = append(eq, i)
			} else {
				neq = append(neq, i)
			}
		}
		if len(neq)%2 == 1 {
			write(f, "NO\n")
			continue
		}

		s1 = make([]byte, n)
		s2 = make([]byte, n)
		for i := 0; i < len(eq); i++ {
			if i < len(eq)/2 {
				s1[eq[i]], s2[eq[i]] = '(', '('
			} else {
				s1[eq[i]], s2[eq[i]] = ')', ')'
			}
		}
		for i := 0; i < len(neq); i += 2 {
			s1[neq[i]], s2[neq[i+1]] = '(', '('
			s1[neq[i+1]], s2[neq[i]] = ')', ')'

		}
		write(f, "YES\n")
		write(f, string(s1), "\n")
		write(f, string(s2), "\n")

	}
}
