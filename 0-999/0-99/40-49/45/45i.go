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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	neg := make([]int, 0)
	all := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i] < 0 {
			neg = append(neg, a[i])
		} else if a[i] > 0 {
			all = append(all, a[i])
		}
	}
	if n == 1 {
		write(f, a[0], "\n")
	} else {
		sort.Ints(neg)
		if len(neg)%2 == 1 {
			neg = neg[:len(neg)-1]
		}
		for i := 0; i < len(neg); i++ {
			all = append(all, neg[i])
		}
		if len(all) == 0 {
			write(f, "0\n")
		} else {
			for i := 0; i < len(all); i++ {
				write(f, all[i], " ")
			}
			write(f, "\n")
		}
	}
}
