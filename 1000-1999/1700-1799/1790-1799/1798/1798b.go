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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var m int
		fmt.Fscan(reader, &m)
		a := make([][]int, m)
		for i := 0; i < m; i++ {
			var n int
			fmt.Fscan(reader, &n)
			a[i] = make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Fscan(reader, &a[i][j])
			}
		}
		p := make([]int, m)
		ok := true
		mem := map[int]bool{}
		for i := m - 1; i >= 0; i-- {
			found := false
			for j := 0; j < len(a[i]); j++ {
				if !found && !mem[a[i][j]] {
					found = true
					p[i] = a[i][j]
				}
				mem[a[i][j]] = true
			}
			if !found {
				ok = false
				break
			}
		}
		if ok {
			for i := 0; i < m; i++ {
				write(f, p[i], " ")
			}
			write(f, "\n")
		} else {
			write(f, "-1\n")
		}
	}
}
