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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		e := make([]int, 100001)
		o := make([]int, 100001)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i%2 == 0 {
				e[a[i]]++
			} else {
				o[a[i]]++
			}
		}
		sort.Ints(a)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				e[a[i]]--
			} else {
				o[a[i]]--
			}
		}
		ok := true
		for i := 1; i <= 100000; i++ {
			if e[i] != 0 || o[i] != 0 {
				ok = false
				break
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
