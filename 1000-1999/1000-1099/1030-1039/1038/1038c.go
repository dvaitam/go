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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	as, bs := int64(0), int64(0)
	i, j := n-1, n-1
	aturn := true
	for i >= 0 || j >= 0 {
		if aturn {
			if i >= 0 && j >= 0 {
				if a[i] > b[j] {
					as += int64(a[i])
					i--
				} else {
					j--
				}
			} else if i >= 0 {
				as += int64(a[i])
				i--
			} else {
				j--
			}
		} else {
			if i >= 0 && j >= 0 {
				if b[j] > a[i] {
					bs += int64(b[j])
					j--
				} else {
					i--
				}
			} else if j >= 0 {
				bs += int64(b[j])
				j--
			} else {
				i--
			}
		}
		aturn = !aturn
	}
	write(f, as-bs, "\n")

}
