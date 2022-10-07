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
func is_correct(n int64) bool {
	s := fmt.Sprint(n)
	for i := 0; i < len(s); i++ {
		if s[i] != '0' && n%int64(s[i]-'0') != 0 {
			return false
		}
	}
	return true

}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int64
		fmt.Fscan(reader, &n)
		for !is_correct(n) {
			n++
		}
		write(f, n, "\n")
	}
}
