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
func main() {
	var n, a, b, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &a, &b, &k)
	ab := a + b
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
		h[i] = h[i] % ab
		if h[i] == 0 {
			h[i] = ab
		}

		if h[i]%a == 0 {
			h[i] = h[i] / a
			h[i]--
		} else {
			h[i] = h[i] / a
		}
	}
	sort.Ints(h)
	ans := 0
	for i := 0; i < n; i++ {
		if h[i] == 0 {
			ans++
		} else if h[i] <= k {
			k -= h[i]
			ans++
		}
	}
	write(f, ans, "\n")
}
