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
	vow := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'n': true}
	ok := true
	for i := 0; i < len(s); i++ {
		if !vow[s[i]] {
			if i == len(s)-1 {
				ok = false
				break
			}
			if i+1 < len(s) && (!vow[s[i+1]] || s[i+1] == 'n') {
				ok = false
				break
			}
		}
	}
	if ok {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}
}
