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

type Dsu struct {
	p []int
}
type Point struct {
	x, y int
}

func (d Dsu) get(x int) int {
	if x == d.p[x] {
		return x
	} else {
		d.p[x] = d.get(d.p[x])
		return d.p[x]
	}
}
func (d Dsu) unite(x int, y int) bool {
	x = d.get(x)
	y = d.get(y)
	if x != y {
		d.p[x] = y
		return true
	}
	return false
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	points := make([]Point, n)
	dsu := Dsu{}
	dsu.p = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dsu.p[i] = i
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].y)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]
			if p1.x == p2.x || p1.y == p2.y {
				dsu.unite(i+1, j+1)
			}
		}
	}
	m := map[int]bool{}
	for i := 1; i <= n; i++ {
		m[dsu.get(i)] = true
	}
	write(f, len(m)-1, "\n")
}
