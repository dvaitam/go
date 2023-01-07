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

func backtrack(b [][]int, m map[int]int, sm int, n int, start int, ans [][]int) {
	if found {
		return
	}
	if start == n*n {
		//	fmt.Println("checking", m, b)
		ok := true
		for i := 0; i < n; i++ {
			tmp := 0
			for j := 0; j < n; j++ {
				tmp += b[i][j]
			}
			if tmp != sm {
				ok = false
				break
			}
		}
		if ok {
			for i := 0; i < n; i++ {
				tmp := 0
				for j := 0; j < n; j++ {
					tmp += b[j][i]
				}
				if tmp != sm {
					ok = false
					break
				}
			}
		}
		if ok {
			tmp := 0
			for i := 0; i < n; i++ {
				tmp += b[i][i]
			}
			if tmp != sm {
				ok = false
			}
			tmp = 0
			for i := 0; i < n; i++ {
				tmp += b[i][n-i-1]
			}
			if tmp != sm {
				ok = false
			}
		}
		if ok {
			found = true
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					ans[i][j] = b[i][j]
				}
			}
		}
		return
	}
	for k := range m {
		if m[k] > 0 {
			b[start/n][start%n] = k
			m[k]--
			if !found {
				backtrack(b, m, sm, n, start+1, ans)
			}
			m[k]++
		}
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n*n)
	b := make([][]int, n)
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
		ans[i] = make([]int, n)
	}
	sm := 0
	m := map[int]int{}
	for i := 0; i < n*n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]]++
		sm += a[i]
	}
	sm = sm / n
	backtrack(b, m, sm, n, 0, ans)
	write(f, sm, "\n")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			write(f, ans[i][j])
			if j == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	}
}
