// 1294A. Collecting Coins
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
		var a, b, c, n int
		fmt.Fscan(reader, &a, &b, &c, &n)
		mx := max(a, max(b, c))
		rm := 3*mx - (a + b + c)
		if n < rm || (n-rm)%3 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
