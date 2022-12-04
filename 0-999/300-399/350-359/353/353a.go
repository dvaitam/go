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
	b := make([]int, n)
	as := 0
	bs := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		as += a[i]
		bs += b[i]
	}
	if as%2 == 0 && bs%2 == 0 {
		write(f, "0\n")
	} else {
		ok := false
		for i := 0; i < n; i++ {
			if (as+b[i]-a[i])%2 == 0 && (bs+a[i]-b[i])%2 == 0 {
				ok = true
				break
			}
		}
		if ok {
			write(f, "1\n")
		} else {
			write(f, "-1\n")
		}
	}

}
