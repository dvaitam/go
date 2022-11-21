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
func heights(a []int, d []int, i int, j int, h int) {
	if j < i {
		return
	} else if i == j {
		d[i] = h
		return
	} else {
		mx := a[i]
		mi := i
		for k := i; k <= j; k++ {
			if a[k] > mx {
				mx = a[k]
				mi = k
			}
		}
		d[mi] = h
		heights(a, d, i, mi-1, h+1)
		heights(a, d, mi+1, j, h+1)
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
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		d := make([]int, n)
		heights(a, d, 0, n-1, 0)
		for i := 0; i < n; i++ {
			write(f, d[i])
			if i == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	}
}
