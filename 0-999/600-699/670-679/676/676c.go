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
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	var s string
	fmt.Fscan(reader, &s)
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		if i == 0 || s[i] != s[i-1] {
			a = append(a, 1)
		} else {
			a[len(a)-1]++
		}
	}
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			c[i] = a[i]
		} else {
			c[i] += c[i-1] + a[i]
		}
	}
	i := 0
	j := 0
	curr := 0
	ans := 0

	for k := 0; k < len(a); k++ {
		ans = max(ans, a[k])
	}

	for j < len(a) && i < len(a) {
		for j < len(a) && curr+a[j] <= k {
			curr += a[j]
			j += 2
		}
		sm := 0
		end := min(len(a)-1, j-1)
		if end >= 0 {
			sm += c[end]
		}
		if i-2 >= 0 {
			sm -= c[i-2]
		}
		extra := 0
		if j < len(a) {
			extra += a[j]
		}
		if i-2 >= 0 {
			extra += a[i-2]
		}
		sm += min(extra, k-curr)
		ans = max(ans, sm)
		curr -= a[i]
		i += 2
	}
	i = 1
	j = 1
	curr = 0
	for j < len(a) && i < len(a) {
		for j < len(a) && curr+a[j] <= k {
			curr += a[j]
			j += 2
		}
		sm := 0
		end := min(len(a)-1, j-1)
		if end >= 0 {
			sm += c[end]
		}
		//write(f, "sum is", sm, end, "\n")
		if i-2 >= 0 {
			sm -= c[i-2]
		}
		//write(f, "sum is", sm, i-2, "\n")

		extra := 0
		if j < len(a) {
			extra += a[j]
		}
		if i-2 >= 0 {
			extra += a[i-2]
		}
		sm += min(extra, k-curr)
		ans = max(ans, sm)
		curr -= a[i]
		i += 2
	}
	write(f, ans, "\n")

}
