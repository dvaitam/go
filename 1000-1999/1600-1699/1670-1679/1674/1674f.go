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
	var n, m, q int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &q)
	a := make([]bool, m*n)
	s := make([]string, n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		for j := 0; j < m; j++ {
			if s[i][j] == '*' {
				a[j*n+i] = true
				count++
			}
		}
	}
	empty := 0
	for i := 0; i < count; i++ {
		if !a[i] {
			empty++
		}
	}
	for i := 0; i < q; i++ {
		//write(f, a, "\n")
		var x, y int
		fmt.Fscan(reader, &x, &y)
		x--
		y--
		pos := y*n + x
		if a[pos] {
			if !a[count-1] {
				empty--
			}
			count--
			if pos < count {
				empty++
			}
		} else {
			count++
			if !a[count-1] {
				empty++
			}
			if pos < count {
				empty--
			}
		}
		a[pos] = !a[pos]
		write(f, empty, "\n")
	}

}
