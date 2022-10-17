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
		var n, curr int
		fmt.Fscan(reader, &n)
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			m[curr]++
		}
		for true {
			rep := -1
			for k, v := range m {
				if v > 1 {
					rep = k
					break
				}
			}
			if rep == -1 || rep == 2048 {
				break
			} else {
				m[rep] -= 2
				m[2*rep] += 1
			}
		}
		if m[2048] > 0 {
			write(f, "YES\n")
		} else {
			write(f, "NO\n")
		}
	}
}
