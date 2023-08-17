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
		var s string
		fmt.Fscan(reader, &n, &s)
		curr := 1
		zeros := make([]int, 0)
		ones := make([]int, 0)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				if len(zeros) > 0 {
					ll := len(zeros) - 1
					a[i] = zeros[ll]
					ones = append(ones, zeros[ll])
					zeros = zeros[:ll]
				} else {
					a[i] = curr
					ones = append(ones, curr)
					curr++
				}
			} else {
				if len(ones) > 0 {
					ll := len(ones) - 1
					a[i] = ones[ll]
					zeros = append(zeros, ones[ll])
					ones = ones[:ll]
				} else {
					a[i] = curr
					zeros = append(zeros, curr)
					curr++
				}
			}
		}
		write(f, curr-1, "\n")
		for i := 0; i < n; i++ {
			write(f, a[i], " ")
		}
		write(f, "\n")
	}
}
