// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m, curr int
		fmt.Fscan(reader, &n, &m)
		bs := make([]byte, m)
		for i := 0; i < m; i++ {
			bs[i] = 'B'
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			first := curr - 1
			other := m - curr
			if first > other {
				first, other = other, first
			}
			if bs[first] == 'B' {
				bs[first] = 'A'
			} else {
				bs[other] = 'A'
			}
		}
		fmt.Println(string(bs))
	}
}
