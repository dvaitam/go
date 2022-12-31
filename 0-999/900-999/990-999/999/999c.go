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
func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s string
	fmt.Fscan(reader, &n, &k, &s)
	m := map[byte]int{}
	for i := 0; i < n; i++ {
		m[s[i]]++
	}
	keys := make([]byte, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for i := 0; i < len(keys); i++ {
		if m[keys[i]] <= k {
			k -= m[keys[i]]
			m[keys[i]] = 0
		} else {
			m[keys[i]] -= k
			k = 0
		}
	}
	used := make([]bool, n)
	for i := n - 1; i >= 0; i-- {
		if m[s[i]] > 0 {
			used[i] = true
			m[s[i]]--
		}
	}
	ans := make([]byte, 0)
	for i := 0; i < n; i++ {
		if used[i] {
			ans = append(ans, s[i])
		}
	}
	write(f, string(ans), "\n")
}
