// 1543A. Exciting Bets
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
		var a, b int64
		fmt.Fscan(reader, &a, &b)
		if a < b {
			a, b = b, a
		}
		if a == b {
			fmt.Println(0, 0)
		} else {
			ans := a - b
			b = b % ans
			if b < ans-b {
				fmt.Println(ans, b)
			} else {
				fmt.Println(ans, ans-b)
			}
		}

	}
}
