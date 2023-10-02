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
	var n, a, x, b, y int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &a, &x, &b, &y)
	ok := false
	for true {
		//write(f, a, x, b, y, "\n")
		if a == b {
			ok = true
		}
		a++
		b--
		if a > n {
			a = 1
		}
		if b < 1 {
			b = n
		}
		if a == b {
			ok = true
		}
		if a == x || b == y {
			break
		}
	}
	if ok {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
