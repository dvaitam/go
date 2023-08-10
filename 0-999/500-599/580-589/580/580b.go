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

type Friend struct {
	m, s int64
}

func main() {
	var n int
	var d int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &d)
	a := make([]Friend, n)
	for i := 0; i < n; i++ {
		var m, s int64
		fmt.Fscan(reader, &m, &s)
		a[i] = Friend{m: m, s: s}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].m < a[j].m })
	count := int64(0)
	ans := int64(0)
	i, j := 0, 0
	for i < n {
		for j < n && a[j].m < a[i].m+d {
			count += a[j].s
			j++
		}
		//	write(f, count, "\n")
		ans = max(ans, count)
		if j == n {
			break
		}
		if i == j {
			count -= a[i].s
			i++
			j++
		} else {
			count -= a[i].s
			i++
		}
	}
	ans = max(ans, count)
	write(f, ans, "\n")
}
