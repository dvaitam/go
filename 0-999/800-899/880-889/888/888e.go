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

var ans int

func get_rem(a []int, i int, n int, m int, rem int) {
	if i == n || ans == m-1 {
		return
	}
	//cache[rem] = true
	ans = max(ans, (rem+a[i])%m)
	ans = max(ans, rem)
	get_rem(a, i+1, n, m, rem)
	get_rem(a, i+1, n, m, (rem+a[i])%m)

}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int64, n)
	write(f, 171673>>10&1, "\n")
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	b := a[:(n+1)/2]
	fa := func(i int) int {
		s := int64(0)
		for j, v := range b {
			//write(f, i, j, i>>j&1, "\n")

			if i>>j&1 > 0 {
				s += v
			}
		}
		return int(s % int64(m))
	}
	x := []int{}
	for i := 0; i < 1<<len(b); i++ {
		s := fa(i)
		ans = max(ans, s)
		x = append(x, s)
	}
	sort.Ints(x)
	b = a[(n+1)/2:]
	for i := 0; i < 1<<len(b); i++ {
		s := fa(i)
		ans = max(ans, s)
		if j := sort.SearchInts(x, m-s); j > 0 {
			ans = max(ans, s+x[j-1])
		}
	}

	write(f, ans, "\n")
}
