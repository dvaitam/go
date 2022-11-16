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
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		mx := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i == 0 {
				mx[i] = a[i]
			} else {
				if mx[i-1] < a[i] {
					mx[i] = a[i]
				} else {
					mx[i] = mx[i-1]
				}
			}
		}
		ans := int64(0)
		i := 0
		for i < n {
			start := mx[i]
			for mx[i] == start {
				if i == 0 {
					i++
				} else {
					if a[i] < a[i-1] {
						ans += int64(a[i-1] - a[i])
					}
					i++
				}
				if i == n {
					break
				}
			}
		}
		write(f, ans, "\n")
	}
}
