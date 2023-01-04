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
type Student struct {
	x, t, i int
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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	s := make([]Student, n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i].t, &s[i].x)
		s[i].i = i
	}
	i := 0
	arrival_time := 0
	k := 1
	for i < n {
		start := i
		end := i + m
		if end > n {
			end = n
		}
		bus_start := s[end-1].t
		bus_start = max(bus_start, arrival_time)
		sort.Slice(s[start:end], func(i1, i2 int) bool { return s[start+i1].x < s[start+i2].x })
		next_k := 1
		last_time := -1
		//write(f, "bus start ", end-1, bus_start, "\n")
		for j := start; j < end; j++ {
			if j == start {
				t[s[j].i] = bus_start + s[j].x
				//	write(f, "begin ", bus_start, s[j].x, "\n")
			} else {
				if s[j].x == s[j-1].x {
					t[s[j].i] = last_time
					next_k++
				} else {
					k = next_k
					t[s[j].i] = last_time + 1 + (k / 2) + s[j].x - s[j-1].x
					next_k = 1
				}
			}
			last_time = t[s[j].i]
		}
		k = next_k
		arrival_time = last_time + 1 + (k / 2) + s[end-1].x
		//write(f, "arrival ", arrival_time, "\n")
		i += m
	}
	for j := 0; j < n; j++ {
		write(f, t[j])
		if j == n-1 {
			write(f, "\n")
		} else {
			write(f, " ")
		}
	}

}
