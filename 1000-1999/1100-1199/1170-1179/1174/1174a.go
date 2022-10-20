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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, 2*n)

	for i := 0; i < 2*n; i++ {
		fmt.Fscan(reader, &a[i])

	}
	sort.Ints(a)
	if a[0] == a[2*n-1] {
		write(f, "-1\n")
	} else {
		for i := 0; i < 2*n; i++ {
			write(f, a[i])
			if i == 2*n {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	}
}
