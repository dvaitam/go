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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, 4*n)
		m := map[int]int{}
		for i := 0; i < 4*n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]]++
		}
		ok := true
		half := make([]int, 0)
		for k, v := range m {
			if v%2 != 0 {
				ok = false
				break
			} else {
				for i := 0; i < v/2; i++ {
					half = append(half, k)
				}
			}
		}
		if ok {
			sort.Ints(half)
			area := half[0] * half[2*n-1]
			for i := 1; i < n; i++ {
				if area != half[i]*half[2*n-i-1] {
					ok = false
					break
				}
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
