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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	even := 0
	odd := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if odd == 1 {
			if a[i]%2 == 1 {
				ans = i + 1
			}
		} else {
			if a[i]%2 == 0 {
				ans = i + 1
			}
		}
	}
	write(f, ans, "\n")

}
