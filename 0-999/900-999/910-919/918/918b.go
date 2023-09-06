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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	mp := map[string]string{}
	for i := 0; i < n; i++ {
		var name, ip string
		fmt.Fscan(reader, &name, &ip)
		mp[ip] = name
	}
	for i := 0; i < m; i++ {
		var command, ip string
		fmt.Fscan(reader, &command, &ip)
		write(f, command, " ", ip, " #", mp[ip[:len(ip)-1]], "\n")
	}
}
