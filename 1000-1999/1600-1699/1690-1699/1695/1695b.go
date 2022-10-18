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

		mi := 1000000000

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] < mi {
				mi = a[i]
			}
		}

		if n%2 == 1 {
			write(f, "Mike\n")
		} else {

			ok := true
			for i := 0; i < n; i++ {
				if a[i]-mi == 0 {
					if i%2 == 0 {
						ok = false
						break
					} else if i%2 == 1 {
						ok = true
						break
					}

				}
			}
			if ok {
				write(f, "Mike\n")
			} else {
				write(f, "Joe\n")
			}
		}
	}
}
