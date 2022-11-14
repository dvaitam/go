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
		var n int
		fmt.Fscan(reader, &n)
		var s string
		fmt.Fscan(reader, &s)
		ans := 0
		started := false
		count := 0
		if s[0] == '1' {
			started = true
			count = 1
		}
		for i := 1; i < n; i++ {
			if s[i] == '1' {
				if started {
					count++
				} else {
					started = true
					count = 1
				}
			} else if started {
				if count > 1 {
					ans += count - 1
				}
				count = 0
				started = false
			}
		}
		if s[n-1] == '1' {
			ans += count - 1
		}
		zans := 0
		count = 0
		if s[0] == '0' {
			started = true
			count = 1
		}
		for i := 1; i < n; i++ {
			if s[i] == '0' {
				if started {
					count++
				} else {
					started = true
					count = 1
				}
			} else if started {
				if count > 1 {
					zans += count - 1
				}
				count = 0
				started = false
			}
		}
		if s[n-1] == '0' {
			zans += count - 1
		}
		write(f, max(ans, zans), "\n")

	}
}
