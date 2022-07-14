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
		var a, b int
		fmt.Fscan(reader, &a, &b)
		c1, c2 := 0, 0
		y := b
		for a|y != y {
			y++
			c1++
		}
		if a != y {
			c1++
		}
		for a|b != b {
			a++
			c2++
		}
		if a != b {
			c2++
		}
		if c2 < c1 {
			c1 = c2
		}
		fmt.Println(c1)

	}
}
