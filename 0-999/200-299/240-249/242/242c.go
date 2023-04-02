// 00
package main

import (
	"bufio"
	"container/list"
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

type Point struct {
	x, y, d int64
}

func main() {
	mod := int64(1000000001)
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var x0, y0, x1, y1 int64
	fmt.Fscan(reader, &x0, &y0, &x1, &y1)
	var n int
	fmt.Fscan(reader, &n)
	m := map[int64]bool{}
	for i := 0; i < n; i++ {
		var r, a, b int64
		fmt.Fscan(reader, &r, &a, &b)
		for i := a; i <= b; i++ {
			m[r*mod+i] = true
		}
	}
	l := list.New()
	l.PushBack(Point{x: x0, y: y0, d: 0})
	visited := map[int64]bool{}
	d := map[int64]int64{}
	visited[x0*mod+y0] = true
	for l.Len() > 0 {
		front := l.Front().Value.(Point)
		l.Remove(l.Front())
		for i := int64(-1); i <= 1; i++ {
			for j := int64(-1); j <= 1; j++ {
				index := (front.x+i)*mod + front.y + j
				if m[index] {
					if !visited[index] {
						l.PushBack(Point{x: front.x + i, y: front.y + j, d: front.d + 1})
						visited[index] = true
						d[index] = front.d + 1
					}
				}
			}
		}

	}
	dest := x1*mod + y1
	if d[dest] > 0 {
		write(f, d[dest], "\n")
	} else {
		write(f, "-1\n")
	}

}
