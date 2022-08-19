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
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		odds := make([]int, 0)
		for i := 0; i < n; i++ {
			curr := s[i] - '0'
			if curr%2 == 1 {
				odds = append(odds, int(curr))
			}
			if len(odds) == 2 {
				break
			}
		}
		if len(odds) == 2 {
			fmt.Println(10*odds[0] + odds[1])
		} else {
			fmt.Println(-1)
		}
	}
}
