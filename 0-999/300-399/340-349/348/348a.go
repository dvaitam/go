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
	a := make([]int64, n)
	mx := int64(0)
	su := int64(0)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		mx = max(mx, a[i])
		su += a[i]
	}
	su = int64(n)*mx - su
	if su >= mx {
		write(f, mx, "\n")
	} else {
		k := (mx - su) / int64(n-1)
		if (mx-su)%int64(n-1) != 0 {
			k++
		}
		write(f, mx+k, "\n")
	}
}
