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

func check(u int, v int) bool {

	if u == v {
		return true
	}
	if u > v {
		return false
	}
	for i := 30; i >= 0; i-- {
		curr := 1 << i
		if curr > v {
			continue
		}
		if curr|v == curr+v && v-curr >= u {
			return check(u, v-curr)
		}
	}
	return false
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		if check(u, v) {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
