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
		var n int
		var h int64
		fmt.Fscan(reader, &n, &h)
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		left := int64(1)
		right := h
		for left < right {
			mid := (left + right) / 2
			tmp := int64(0)
			for i := 0; i < n; i++ {
				if i == n-1 {
					tmp += mid
				} else {
					tmp += min(a[i+1]-a[i], mid)
				}
			}
			if tmp < h {
				left = mid + 1
			} else if tmp > h {
				right = mid
			} else {
				break
			}
		}
		write(f, (left+right)/2, "\n")

	}
}
