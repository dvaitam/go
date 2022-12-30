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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	mia := -1
	mxa := 0
	mib := -1
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if mia == -1 {
			mia = a[i]
		} else {
			mia = min(mia, a[i])
		}
		mxa = max(mxa, a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
		if mib == -1 {
			mib = b[i]
		} else {
			mib = min(mib, b[i])
		}
	}
	ans := max(mxa, 2*mia)
	if ans < mib {
		write(f, ans, "\n")
	} else {
		write(f, "-1\n")
	}

}
