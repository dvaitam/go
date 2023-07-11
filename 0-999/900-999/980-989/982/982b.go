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

type Row struct {
	w, i int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	ws := make([]Row, n)
	for i := 0; i < n; i++ {
		var w int
		fmt.Fscan(reader, &w)
		ws[i] = Row{i: i + 1, w: w}
	}
	sort.Slice(ws, func(i, j int) bool { return ws[i].w < ws[j].w })
	var s string
	fmt.Fscan(reader, &s)
	curr := make([]Row, n)
	index := -1
	j := 0
	for i := 0; i < 2*n; i++ {
		if s[i] == '0' {
			write(f, ws[j].i, " ")
			index++
			curr[index] = ws[j]
			j++
		} else {
			write(f, curr[index].i, " ")
			index--
		}
	}
	write(f, "\n")

}
