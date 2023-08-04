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
	m := map[byte]bool{}
	mark := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'J' {
			m[s[i]] = true
		}
		if s[i] == '?' {
			mark++
		}
	}
	ans := 1
	if s[0] >= 'A' && s[0] <= 'J' {
		ans = ans * 9
		curr := 9
		for i := 0; i < len(m)-1; i++ {
			if curr == 0 {
				break
			}
			ans = ans * curr
			curr--
		}
	} else {
		curr := 10
		for i := 0; i < len(m); i++ {
			if curr == 0 {
				break
			}
			ans = ans * curr
			curr--
		}
	}
	if s[0] == '?' {
		ans = ans * 9
		mark--
	}
	write(f, ans)
	for i := 0; i < mark; i++ {
		write(f, "0")
	}
	write(f, "\n")
}
