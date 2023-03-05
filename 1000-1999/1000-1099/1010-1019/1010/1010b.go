// 00
package main

import (
	"bufio"
	"fmt"
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
	var m, n, r int
	fmt.Scan(&m, &n)
	a := make([]bool, n)
	found := false
	for i := 0; i < n; i++ {
		fmt.Println(m)
		fmt.Scan(&r)
		if r == 1 {
			a[i] = true
		} else if r == 0 {
			found = true
			break
		}
	}
	if !found {
		left, right := 1, m
		i := 0
		for left <= right {
			truth := a[i]
			mid := (left + right) / 2
			fmt.Println(mid)
			fmt.Scan(&r)
			if !truth {
				r = -r
			}
			if r == 1 {
				right = mid
			} else if r == -1 {
				left = mid
			} else {
				break
			}
			i++
			i = i % n
		}

	}

}
