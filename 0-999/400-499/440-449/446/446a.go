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
type Segment struct {
	i, j, length int
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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	segments := make([]Segment, 0)
	start := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if i > 0 {
			if a[i] <= a[i-1] {
				segments = append(segments, Segment{i: start, j: i - 1, length: i - start})
				start = i
			}
		}
	}
	segments = append(segments, Segment{i: start, j: n - 1, length: n - start})
	ans := 0
	m := len(segments)
	for i := 0; i < m; i++ {
		segment := segments[i]
		ans = max(ans, segment.length)
		if i > 0 {
			if segment.length == 1 {
				ans = max(ans, 1+segments[i-1].length)
			} else {
				if a[segment.i+1] > a[segments[i-1].j]+1 {
					ans = max(ans, segment.length+segments[i-1].length)
				} else {
					ans = max(ans, 1+segments[i-1].length)
					ans = max(ans, 1+segment.length)
				}
				if segments[i-1].length > 1 && a[segment.i] > a[segments[i-1].j-1]+1 {
					ans = max(ans, segment.length+segments[i-1].length)
				}
			}
		}
	}
	write(f, ans, "\n")
}
