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
		var s string
		fmt.Fscan(reader, &n, &s)
		b := []byte(s)
		zeros := 0
		ones := 0
		for i := 0; i < n; i++ {
			if b[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		bs := []byte(s)
		for i := 0; i < n; i++ {
			if zeros > 0 {
				bs[i] = '0'
				zeros--
			} else {
				bs[i] = '1'
			}
		}

		k := make([]int, 0)
		for i := 0; i < n; i++ {
			if b[i] != bs[i] {
				k = append(k, i+1)
			}
		}
		if len(k) > 0 {
			write(f, 1, "\n")
			write(f, len(k))
			for i := 0; i < len(k); i++ {
				write(f, " ", k[i])
			}
			write(f, "\n")
		} else {
			write(f, 0, "\n")
		}

	}
}
