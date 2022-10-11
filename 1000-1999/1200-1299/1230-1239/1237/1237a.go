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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	net := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		net += a[i] / 2
	}
	for i := 0; i < n; i++ {
		modified := false

		if a[i]%2 != 0 {
			if net > 0 {
				if a[i] < 0 {
					a[i] = (a[i] / 2) - 1
					net--
					modified = true
				}
			} else if net < 0 {
				if a[i] > 0 {
					a[i] = (a[i] / 2) + 1
					net++
					modified = true
				}
			}
		}
		if !modified {
			a[i] = a[i] / 2
		}
	}
	for i := 0; i < n; i++ {
		write(f, a[i], "\n")
	}

}
