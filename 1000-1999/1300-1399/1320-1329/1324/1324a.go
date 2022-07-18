// 1324A. Yet Another Tetris Problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		mx := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			mx = max(mx, a[i])
		}
		ok := true
		for i := 0; i < n; i++ {
			if (mx-a[i])%2 == 1 {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
