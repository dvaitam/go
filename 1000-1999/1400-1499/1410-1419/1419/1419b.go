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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var x int64
		fmt.Fscan(reader, &x)
		ans := 0
		sm := int64(0)
		curr := int64(0)
		for sm+curr*(curr+1)+(curr+1)*(curr+1) <= x {
			sm += curr*(curr+1) + (curr+1)*(curr+1)
			ans++
			curr = 2*curr + 1
		}
		write(f, ans, "\n")

	}
}
