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
	var n int
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &s)
	a := make([]int, 26)
	for i := 0; i < n; i++ {
		a[s[i]-'a']++
	}
	ok := false
	for i := 0; i < 26; i++ {
		if a[i] > 1 {
			ok = true
			break
		}
	}
	if ok || n == 1 {
		write(f, "Yes\n")
	} else {
		write(f, "No\n")
	}
}
