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

const MAXN = 1000000

var n, w int
var a, b, tx [MAXN]int

func fail() {
	tx[0] = -1
	pos := 1
	cnd := 0
	for pos <= w {
		if cnd == 0 || b[pos-1] == b[cnd-1] {
			tx[pos] = cnd
			pos++
			cnd++
		} else {
			cnd = tx[cnd]
		}
	}
}
func matches() int {
	res := 0
	pos := 0
	cnd := 0
	for pos < n {
		if cnd == -1 || (cnd < w && b[cnd] == a[pos]) {
			cnd++
			pos++
			if cnd == w {
				res++
			}
		} else {
			cnd = tx[cnd]
		}
	}
	return res
}

func main() {
	//	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	//fmt.Fscan(reader, &T)
	//for t := 1; t <= T; t++ {
	fmt.Fscan(reader, &n, &w)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < w; i++ {
		fmt.Fscan(reader, &b[i])
	}
	if w == 1 {
		write(f, n, "\n")
		return
	}
	n--
	for i := 0; i < n; i++ {
		a[i] -= a[i+1]
	}
	w--
	for i := 0; i < w; i++ {
		b[i] -= b[i+1]
	}

	fail()
	write(f, matches(), "\n")

	// }
}
