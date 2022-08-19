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
		var n, x int
		fmt.Fscan(reader, &n, &x)
		a := make([]int, n)
		m := make(map[int]bool)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = true
		}
		curr := 1

		for x > 0 {
			if !m[curr] {
				x--
			}
			curr++
		}
		for m[curr] {
			curr++
		}
		fmt.Println(curr - 1)
	}
}
