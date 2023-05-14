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
	e, l, r int
}

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	sorted := make([]Segment, 0)
	unsorted := make([]Segment, 0)
	//tl, tr := -1, -1
	for i := 0; i < m; i++ {
		var t, l, r int
		fmt.Fscan(reader, &t, &l, &r)
		// if i == 613 {
		// 	tl, tr = l, r
		// }
		// if n == 1000 {
		// 	if r == 1000 {
		// 		write(f, t, l, r, "\n")
		// 	}
		// }
		if t == 1 {
			sorted = append(sorted, Segment{l: l - 1, r: r - 1})
		} else {
			unsorted = append(unsorted, Segment{l: l - 1, r: r - 1})
		}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].l < sorted[j].l || sorted[i].l == sorted[j].l && sorted[i].r < sorted[j].r
	})
	new_segments := make([]Segment, 0)
	for i := 0; i < len(sorted); i++ {
		if len(new_segments) == 0 {
			new_segments = append(new_segments, sorted[i])
		} else {
			if sorted[i].l <= new_segments[len(new_segments)-1].r {
				new_segments[len(new_segments)-1].r = max(new_segments[len(new_segments)-1].r, sorted[i].r)
			} else {
				new_segments = append(new_segments, sorted[i])
			}
		}
	}
	ok := true
	// for i := 0; i < len(unsorted); i++ {
	// 	for j := 0; j < len(new_segments); j++ {
	// 		mi, mx := max(new_segments[j].l, unsorted[i].l), min(new_segments[j].r, unsorted[i].r)
	// 		if mx < mi {
	// 			continue
	// 		}
	// 		if unsorted[i].l >= new_segments[i].l && unsorted[i].r <= new_segments[i].r {
	// 			ok = false
	// 			break
	// 		}
	// 		if mi > unsorted[i].l {
	// 			unsorted[i].e = -1
	// 			break
	// 		} else if mx < unsorted[i].r {
	// 			unsorted[i].e = 1
	// 			break
	// 		}
	// 	}
	// 	if !ok {
	// 		break
	// 	}
	// }
	// if !ok {
	// 	write(f, "NO\n")
	// 	return
	// }
	//sr := make([]bool, n)
	// for i := 0; i < n; i++ {
	// 	sr[i] = true
	// }
	ans := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	ans[i] = i + 1
	// }
	// for i := 0; i < len(unsorted); i++ {
	// 	l, r, _ := unsorted[i].l, unsorted[i].r, unsorted[i].e
	// 	for j := l; j <= r; j++ {
	// 		sr[j] = false
	// 	}
	// }
	// write(f, sr, "\n")

	curr := 2000
	prev := -1
	for i := 0; i < len(new_segments); i++ {
		l, r, _ := new_segments[i].l, new_segments[i].r, new_segments[i].e
		for j := prev + 1; j < l; j++ {
			ans[j] = curr
			curr--
		}
		curr--
		for j := l; j <= r; j++ {
			ans[j] = curr
		}
		curr--
		prev = r
	}
	//	write(f, new_segments, ans, "\n")
	curr--
	//curr = 1000
	for i := 0; i < n; i++ {
		if ans[i] == 0 {
			ans[i] = curr
			curr--
		}
	}
	for i := 0; i < len(unsorted); i++ {
		l, r, _ := unsorted[i].l, unsorted[i].r, unsorted[i].e
		good := false
		for j := l; j <= r; j++ {
			if j+1 <= r {
				if ans[j] <= ans[j+1] {
					continue
				} else {
					good = true
					break
				}
			}
		}
		if !good {
			ok = false
			break
		}
	}
	// if m == 1000 {
	// 	ok = true
	// }
	//write(f, ans, "\n")
	if !ok {
		write(f, "NO\n")
		return
	}

	write(f, "YES\n")

	for i := 0; i < n; i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
