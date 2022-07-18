// 00
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
func min(a int, b int) int {
	if a < b {
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
		h := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &h[i])
		}
		ans := 0
		for i := 1; i < n-1; i += 2 {
			mx := max(h[i-1], h[i+1])
			if mx >= h[i] {
				ans += (mx - h[i] + 1)
			}
		}
		if n%2 == 1 {
			fmt.Println(ans)
		} else {
			ans1 := 0
			for i := 2; i < n-1; i += 2 {
				mx := max(h[i-1], h[i+1])
				if mx >= h[i] {
					ans1 += (mx - h[i] + 1)
				}
			}
			fmt.Println(min(ans1, ans))

		}

	}
}
