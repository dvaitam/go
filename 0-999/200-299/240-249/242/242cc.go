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

type Point struct {
	x, y int64
}
type Segment struct {
	a, b int64
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
	segments := map[int64][]Segment{}
	for i := 0; i < n; i++ {
		var r, a, b int64
		fmt.Fscan(reader, &r, &a, &b)
		segments[r] = append(segments[r], Segment{a: a, b: b})
	}
	segments[x0] = append(segments[x0], Segment{a: y0, b: y0})
	segments[x1] = append(segments[x1], Segment{a: y1, b: y1})

	for k := range segments {
		sort.Slice(segments[k], func(i, j int) bool { return segments[k][i].a < segments[k][j].a })
	}
	rsegments := map[int64][]Segment{}
	for k := range segments {
		new_segments := make([]Segment, 0)

		for i := 0; i < len(segments[k]); i++ {
			nl := len(new_segments)

			if nl == 0 || segments[k][i].a > new_segments[nl-1].b {
				new_segments = append(new_segments, segments[k][i])
			} else {
				new_segments[nl].b = max(segments[k][i].b, new_segments[nl].b)
			}
		}
		rsegments[k] = new_segments
	}
	adj := map[int64]map[int64]bool{}
	for k, s := range rsegments {
		pl := len(rsegments[k-1])
		if pl > 0 {
			for ri := 0; ri < len(s); ri++ {
				for sp := s[ri].a - 1; sp <= s[ri].a+1; sp++ {
					j := sort.Search(pl, func(i int) bool {
						return rsegments[k-1][i].a >= sp
					})
					if j < pl && sp >= segments[k-1][j].a {
						source, dest := k*mod+s[ri].a, (k-1)*mod+sp
						if adj[source] == nil {
							adj[source] = map[int64]bool{}
						}
						if adj[dest] == nil {
							adj[dest] = map[int64]bool{}
						}
						adj[source][dest] = true
						adj[dest][source] = true
					}
				}
				for sp := s[ri].b - 1; sp <= s[ri].b+1; sp++ {
					j := sort.Search(pl, func(i int) bool {
						return rsegments[k-1][i].a >= sp
					})
					if j < pl && sp >= segments[k-1][j].a {
						source, dest := k*mod+s[ri].a, (k-1)*mod+sp
						if adj[source] == nil {
							adj[source] = map[int64]bool{}
						}
						if adj[dest] == nil {
							adj[dest] = map[int64]bool{}
						}
						adj[source][dest] = true
						adj[dest][source] = true
					}
				}
			}
		}
	}

}
