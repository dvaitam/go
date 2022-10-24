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
	var s string
	fmt.Fscan(reader, &s)
	ans := 0
	count := 0
	start := false
	m := map[byte]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'Y': true}
	for i := 0; i < len(s); i++ {
		if start && !m[s[i]] {
			count++
		} else if start {
			ans = max(ans, count)
			start = false
			count = 0
		} else if !m[s[i]] {
			start = true
			count = 1
		}
	}
	if start {
		ans = max(ans, count)
	}
	ans++
	write(f, ans, "\n")

}
