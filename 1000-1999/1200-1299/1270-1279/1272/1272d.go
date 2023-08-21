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

type Segment struct {
	l, r int
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
	segments := make([]Segment, 1)
	segments[0] = Segment{l: 0, r: 0}
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			segments[len(segments)-1].r++
		} else {
			segments = append(segments, Segment{l: i, r: i})
		}
	}
	ans := 1
	for i := 0; i < len(segments); i++ {
		ans = max(ans, segments[i].r-segments[i].l+1)
		if i+1 < len(segments) {
			if segments[i].r-segments[i].l > 0 {
				if a[segments[i+1].l] > a[segments[i].r-1] {
					ans = max(ans, segments[i+1].r-segments[i].l)
				}
			}

			if segments[i+1].r-segments[i+1].l > 0 {
				if a[segments[i+1].l+1] > a[segments[i].r] {
					ans = max(ans, segments[i+1].r-segments[i].l)
				}
			} else {
				if i+2 < len(segments) {
					if a[segments[i+2].l] > a[segments[i].r] {
						ans = max(ans, segments[i+2].r-segments[i].l)
					}
				}
			}
		}
	}
	write(f, ans, "\n")
}
