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
	var d int
	var of, wm string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &d, &of, &wm)
	if wm == "week" {
		if d == 5 || d == 6 {
			write(f, "53\n")
		} else {
			write(f, "52\n")
		}
	} else {
		if d < 30 {
			write(f, "12\n")
		} else if d == 30 {
			write(f, "11\n")
		} else {
			write(f, "7\n")
		}
	}
}
