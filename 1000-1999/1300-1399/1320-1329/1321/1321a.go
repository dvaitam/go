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
	r := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &r[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	both := 0
	rwin := 0
	bwin := 0
	for i := 0; i < n; i++ {
		if r[i] == 1 && b[i] == 0 {
			rwin++
		} else if r[i] == 0 && b[i] == 1 {
			bwin++
		} else {
			both++
		}
	}
	if rwin == 0 {
		write(f, "-1\n")
	} else {
		bwin++
		ans := bwin / rwin
		if bwin%rwin != 0 {
			ans++
		}
		write(f, ans, "\n")
	}
}
