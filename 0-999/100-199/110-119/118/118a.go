// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	vow := map[byte]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'Y': true, 'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}
	for i := 0; i < len(s); i++ {
		if !vow[s[i]] {
			write(f, ".", strings.ToLower(s[i:i+1]))
		}
	}
	write(f, "\n")
}
