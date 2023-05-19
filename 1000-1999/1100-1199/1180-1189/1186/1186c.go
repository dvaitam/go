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
	var s, t string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s, &t)
	m, n := true, true
	same, diff := true, true
	for i := 0; i < len(t); i++ {
		if s[i] == t[i] {
			m = !m
		} else {
			n = !n
		}
		if i > 0 {
			if t[i] == t[i-1] {
				same = !same
			} else {
				diff = !diff
			}
		}
	}
	ans := 0
	if n {
		ans++
	}
	//write(f, "ans ", ans, "\n")
	for i := len(t); i < len(s); i++ {
		if s[i-len(t)] == t[0] {
			m = !m
		} else {
			n = !n
		}
		if !diff {
			m = !m
			n = !n
		}
		if s[i] == t[len(t)-1] {
			m = !m
		} else {
			n = !n
		}
		//write(f, m, n, "\n")

		if n {
			ans++
		}
	}
	write(f, ans, "\n")

}
