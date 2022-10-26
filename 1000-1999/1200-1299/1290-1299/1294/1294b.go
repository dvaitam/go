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

type Point struct {
	x, y int
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
func write_string(f *bufio.Writer, x1 int, y1 int, x2 int, y2 int) {
	for x1 < x2 || y1 < y2 {
		if x1 < x2 {
			write(f, "R")
			x1++
		} else {
			write(f, "U")
			y1++
		}
	}
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		p := make([]Point, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i].x, &p[i].y)
		}
		sort.Slice(p, func(i, j int) bool { return p[i].x <= p[j].x && p[i].y <= p[j].y })
		ok := true
		for i := 1; i < n; i++ {
			if p[i].x < p[i-1].x || p[i].y < p[i-1].y {
				ok = false
				break
			}
		}
		if ok {
			write(f, "YES\n")
			for i := 0; i < n; i++ {
				if i == 0 {
					write_string(f, 0, 0, p[i].x, p[i].y)
				} else {
					write_string(f, p[i-1].x, p[i-1].y, p[i].x, p[i].y)
				}
			}
			write(f, "\n")
		} else {
			write(f, "NO\n")
		}

	}
}
