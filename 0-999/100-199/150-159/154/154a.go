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
	var k int
	fmt.Fscan(reader, &k)
	m := map[byte]byte{}
	for i := 0; i < k; i++ {
		var t string
		fmt.Fscan(reader, &t)
		m[t[0]] = t[1]
		m[t[1]] = t[0]
	}
	i := 0
	ans := 0
	for i < len(s) {
		if m[s[i]] == 0 {
			i++
		}else{
			start := s[i]
			one := 0
			other := 0
			for i < len(s) && (s[i] == start || s[i] == m[start]) {
				if s[i] == start {
					one++
				}else{
					other++
				}
				i++
			}
			ans += min(one, other)
		}
	}
	write(f, ans, "\n")

}
