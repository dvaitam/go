// 1551A. Polycarp and Coins
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
		fmt.Fscan(reader, &n)
		c1, c2 := n/3, n/3
		if n%3 == 1 {
			c1++
		} else if n%3 == 2 {
			c2++
		}
		fmt.Println(c1, c2)

	}
}
