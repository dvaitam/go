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
	l, r int
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
		uniq := make([]int, 0)
		unim := map[int]bool{}
		starting := map[int]int{}
		ending := map[int]int{}
		segments := make([]Segment, n)
		for i := 0; i < n; i++ {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			segments[i] = Segment{l, r}
			starting[l]++
			ending[r]++
			if !unim[l] {
				uniq = append(uniq, l)
				unim[l] = true
			}
			if !unim[r] {
				uniq = append(uniq, r)
				unim[r] = true
			}
		}
		sort.Ints(uniq)
		ans := n
		started := 0
		ended := 0
		index := map[int]int{}
		seen := make([]int, len(uniq))
		unseen := make([]int, len(uniq))
		for i := 0; i < len(uniq); i++ {
			started += starting[uniq[i]]
			ended += ending[uniq[i]]
			seen[i] = started
			unseen[i] = ended
			index[uniq[i]] = i
		}
		//write(f, seen, uniq, "\n")
		for i := 0; i < n; i++ {
			s := seen[index[segments[i].r]]
			if index[segments[i].l] > 0 {
				s -= unseen[index[segments[i].l]-1]
			}
			ans = min(ans, n-s)
		}
		write(f, ans, "\n")

	}
}
