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
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		ma := map[int]int{}
		mb := map[int]int{}
		curr := a[0]
		count := 1
		for i := 1; i < n; i++ {
			if a[i] == curr {
				count++
			} else {
				ma[curr] = max(ma[curr], count)
				count = 1
				curr = a[i]
			}
		}
		ma[curr] = max(ma[curr], count)
		curr = b[0]
		count = 1
		for i := 1; i < n; i++ {
			if b[i] == curr {
				count++
			} else {
				mb[curr] = max(mb[curr], count)
				count = 1
				curr = b[i]
			}
		}
		mb[curr] = max(mb[curr], count)
		ans := 0
		for i := 0; i < n; i++ {
			ans = max(ans, ma[a[i]]+mb[a[i]])
			ans = max(ans, ma[b[i]]+mb[b[i]])
		}
		write(f, ans, "\n")

	}
}
