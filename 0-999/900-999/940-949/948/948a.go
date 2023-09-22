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
	var r, c int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &r, &c)
	s := make([]string, r)
	sb := make([][]byte, r)
	for i := 0; i < r; i++ {
		fmt.Fscan(reader, &s[i])
		sb[i] = []byte(s[i])
		for j := 0; j < c; j++ {
			if sb[i][j] == '.' {
				sb[i][j] = 'D'
			}
		}
	}
	ok := true
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if sb[i][j] == 'S' {
				if i+1 < r && sb[i+1][j] == 'W' {
					ok = false
				}
				if i-1 >= 0 && sb[i-1][j] == 'W' {
					ok = false
				}
				if j+1 < c && sb[i][j+1] == 'W' {
					ok = false
				}
				if j-1 >= 0 && sb[i][j-1] == 'W' {
					ok = false
				}
			}
		}
	}
	if ok {
		write(f, "Yes\n")
		for i := 0; i < r; i++ {
			write(f, string(sb[i]), "\n")
		}
	} else {
		write(f, "No\n")
	}
}
