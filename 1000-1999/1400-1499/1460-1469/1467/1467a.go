// 1467A. Wizard of Orz
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
		curr := 9
		inc := false
		for i := 0; i < n; i++ {
			fmt.Print(curr)
			if inc {
				curr++
				curr = curr % 10
			} else {
				curr--
			}
			if curr == 8 && !inc {
				inc = true
			}
		}
		fmt.Println()
	}
}
