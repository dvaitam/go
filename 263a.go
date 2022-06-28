// 282A - Bit++
package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -1 * n
	} else {
		return n
	}
}

func main() {
	var curr int

	ans := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&curr)
			if curr == 1 {
				ans += abs(2 - i)
				ans += abs(2 - j)
			}

		}

	}
	fmt.Println(ans)
}
