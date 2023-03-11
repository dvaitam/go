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

type Edge struct {
	a, b int
}

func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	d := make([]int, n)
	m := map[int][]int{}
	mk := map[int]int{}
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
		if m[d[i]] == nil {
			m[d[i]] = make([]int, 0)
		}
		m[d[i]] = append(m[d[i]], i+1)
		mk[i+1] = k
		mx = max(mx, d[i])
	}
	ok := true
	start := 0
	ans := make([]Edge, 0)
	if len(m[0]) > 1 || len(m[0]) == 0 {
		ok = false
	}
	for ok && len(m[start]) > 0 {

		curr := 0
		for _, b := range m[start+1] {
			if mk[b] == 0 {
				ok = false
				break
			}
			for curr < len(m[start]) && mk[m[start][curr]] == 0 {
				curr++
			}
			if curr == len(m[start]) {
				ok = false
				break
			}
			ans = append(ans, Edge{a: m[start][curr], b: b})
			mk[m[start][curr]]--
			mk[b]--
		}
		start++
	}
	if start != mx+1 {
		ok = false
	}
	if ok {
		write(f, len(ans), "\n")
		for i := 0; i < len(ans); i++ {
			write(f, ans[i].a, ans[i].b, "\n")
		}
	} else {
		write(f, "-1\n")
	}
}
