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
		imp := true
		ans := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i > 0 && i < n-1 {
				if a[i] != 1 {
					imp = false

				}
				ans += a[i] / 2
				if a[i]%2 == 1 {
					ans++
				}
			}

		}
		if n == 3 && a[1]%2 == 1 {
			imp = true
		}
		if imp {
			write(f, "-1\n")
		} else {
			write(f, ans, "\n")
		}
	}
}
