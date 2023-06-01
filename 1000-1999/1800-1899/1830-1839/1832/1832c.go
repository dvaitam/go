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
		b := make([]int, 1)
		b[0] = a[0]
		for i := 1; i < n; i++ {
			if a[i] != b[len(b)-1] {
				if len(b) == 1 {
					b = append(b, a[i])
				} else {
					if b[len(b)-1] > b[len(b)-2] {
						if a[i] < b[len(b)-1] {
							b = append(b, a[i])
						} else {
							b[len(b)-1] = a[i]
						}
					} else {
						if a[i] < b[len(b)-1] {
							b[len(b)-1] = a[i]
						} else {
							b = append(b, a[i])
						}
					}
				}
			}
		}
		//write(f, b, "\n")
		write(f, len(b), "\n")
	}
}
