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
		fmt.Fscan(reader, &n)
		p := make([]int, n)
		index := make([]int, n+1)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
			index[p[i]] = i
		}
		ok := true
		look_for := 1
		for look_for < n {
			for j := index[look_for]; j < n; j++ {
				if p[j] == look_for {
					look_for++
				} else if p[j] < look_for {
					break
				} else {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			write(f, "Yes\n")
		} else {
			write(f, "No\n")
		}

	}
}
