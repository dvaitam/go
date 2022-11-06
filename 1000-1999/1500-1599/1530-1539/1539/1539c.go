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
	var n, k, x int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	gaps := make([]int, 0)
	for i := 1; i < n; i++ {
		gap := a[i] - a[i-1]
		if gap > x {
			g := gap / x
			if gap%x == 0 {
				g--
			}
			gaps = append(gaps, g)
		}
	}
	sort.Ints(gaps)
	rm := 0
	for i := 0; i < len(gaps); i++ {
		if gaps[i] <= k {
			k -= gaps[i]
			rm++
		} else {
			break
		}
	}
	write(f, len(gaps)-rm+1, "\n")
}
