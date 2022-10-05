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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s, t string
	fmt.Fscan(reader, &s, &t)
	ok := true
	if len(s) != len(t) {
		ok = false
	}
	if ok {
		m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
		for i := 0; i < len(s); i++ {
			if m[s[i]] != m[t[i]] {
				ok = false
				break
			}
		}
	}

	if ok {
		write(f, "Yes\n")
	} else {
		write(f, "No\n")
	}
}
