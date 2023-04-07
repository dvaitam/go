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

type Pair struct {
	r, i int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, s1, s2 int
		fmt.Fscan(reader, &n, &s1, &s2)
		r := make([]Pair, n)
		for i := 0; i < n; i++ {
			var rt int
			fmt.Fscan(reader, &rt)
			r[i] = Pair{r: rt, i: i + 1}
		}
		sort.Slice(r, func(i, j int) bool { return r[i].r > r[j].r })
		start1, start2 := 0, 0
		a := make([]int, 0)
		b := make([]int, 0)
		for i := 0; i < n; i++ {
			if start1+s1 <= start2+s2 {
				a = append(a, r[i].i)
				start1 += s1
			} else {
				b = append(b, r[i].i)
				start2 += s2
			}
		}
		write(f, len(a))
		for i := 0; i < len(a); i++ {
			write(f, " ", a[i])
		}
		write(f, "\n")
		write(f, len(b))
		for i := 0; i < len(b); i++ {
			write(f, " ", b[i])
		}
		write(f, "\n")
	}
}
