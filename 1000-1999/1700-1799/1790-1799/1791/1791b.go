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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		x, y := 0, 0
		yes := false
		for i := 0; i < n; i++ {
			if s[i] == 'L' {
				x--
			}else if s[i] == 'R' {
				x++
			}else if s[i] == 'U' {
				y++
			}else{
				y--
			}
			if x == 1 && y == 1 {
				yes = true
			}
		}
		if yes {
			write(f, "YES\n")
		}else{
			write(f, "NO\n")
		}
	}
}
