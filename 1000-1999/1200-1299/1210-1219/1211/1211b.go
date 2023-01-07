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
	mx := int64(-1)
	mxi := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if mx == -1 {
			mx = a[i]
			mxi = 0
		} else {
			if a[i] >= mx {
				mx = a[i]
				mxi = i
			}
		}
	}
	ans := (mx-1)*int64(n) + int64(mxi+1)
	write(f, ans, "\n")

}
