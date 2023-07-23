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
	curr := 0
	ok := true
	for i := 0; i < n; i++ {
		var d int
		var s string
		fmt.Fscan(reader, &d, &s)
		if curr == 0 && s != "South" || curr == 20000 && s != "North" {
			ok = false
		}
		if s == "South" {
			curr += d
		}
		if s == "North" {
			curr -= d
		}
		if curr > 20000 || curr < 0 {
			ok = false
		}
	}
	if curr != 0 {
		ok = false
	}
	if ok {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
