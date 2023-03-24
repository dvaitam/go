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
		var s string
		fmt.Fscan(reader, &s)
		m := map[byte]int{}
		for i := 0; i < 4; i++ {
			m[s[i]]++
		}
		if len(m) == 1 {
			write(f, "-1\n")
		} else {
			mx := 0
			for _, v := range m {
				mx = max(mx, v)
			}
			if mx == 3 {
				write(f, "6\n")
			} else {
				write(f, "4\n")
			}
		}
	}
}
