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
		var x, y int64
		fmt.Fscan(reader, &x, &y)
		if x > y {
			write(f, x+y, "\n")
		} else if x == y || y%x == 0 {
			write(f, x, "\n")
		} else {
			if y-x < x {
				write(f, (x+y)/2, "\n")
			} else {
				trim := y - (y % x)
				write(f, (y+trim)/2, "\n")
			}
		}
	}
}
