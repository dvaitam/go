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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	sb := []byte(s)
	for i := 1; i < len(sb); i++ {
		if sb[i] == sb[i-1] {
			for j := byte(0); j < 26; j++ {
				if sb[i-1] == j+'a' {
					continue
				}
				if i+1 < len(sb) && sb[i+1] == j+'a' {
					continue
				}
				sb[i] = j + 'a'
				break
			}
		}
	}
	write(f, string(sb), "\n")
}
