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
	var s string
	fmt.Fscan(reader, &s)
	m := make(map[byte]int)
	if len(s) > 26 {
		write(f, "-1\n")
	} else {
		for i := 0; i < n; i++ {
			m[s[i]]++
		}
		write(f, n-len(m), "\n")
	}
}
