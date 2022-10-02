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
	start := 1
	end := n * n
	begin := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if begin {
				write(f, start)
				start++
			} else {
				write(f, end)
				end--
			}
			begin = !begin
			if j == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}
	}
}
