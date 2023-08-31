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

type Segment struct {
	l, r, extra int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	segments := make([]Segment, 0)
	for i := 0; i < n; i++ {
		if i == 0 {
			segments = append(segments, Segment{a[i], a[i], 0})
		} else {
			if a[i] == a[i-1] {
				segments[len(segments)-1].extra++
			} else if a[i] == a[i-1]+1 {
				segments[len(segments)-1].r++
			} else {
				segments = append(segments, Segment{a[i], a[i], 0})
			}
		}
	}
	ans := 0
	added := map[int]bool{}
	for i := 0; i < len(segments); i++ {
		if segments[i].extra == 0 {
			ans += segments[i].r - segments[i].l + 1
		} else if segments[i].extra == 1 {
			ans += segments[i].r - segments[i].l + 1
			if segments[i].l > 1 && !added[segments[i].l-1] {
				ans++
				added[segments[i].l-1] = true
			} else {
				ans++
				added[segments[i].r+1] = true
			}
		} else {
			ans += segments[i].r - segments[i].l + 2
			added[segments[i].r+1] = true
			if segments[i].l > 1 && !added[segments[i].l-1] {
				ans++
				added[segments[i].l-1] = true
			}
		}
	}
	write(f, ans, "\n")

}
