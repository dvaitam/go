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
		var a, b string
		fmt.Fscan(reader, &n, &a, &b)
		zeros := make([]int, n)
		ones := make([]int, n)
		if a[0] == '1' {
			ones[0] = 1
		} else {
			zeros[0] = 1
		}
		for i := 1; i < n; i++ {
			if a[i] == '0' {
				zeros[i] = zeros[i-1] + 1
				ones[i] = ones[i-1]
			} else {
				zeros[i] = zeros[i-1]
				ones[i] = ones[i-1] + 1
			}
		}
		ok := true
		flip := false
		for i := n - 1; i >= 0; i-- {
			if (!flip && a[i] == b[i]) || (flip && a[i] != b[i]) {
				continue
			} else if zeros[i] == ones[i] {
				flip = !flip
			} else {
				ok = false
				break
			}
		}
		if ok {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}

	}
}
