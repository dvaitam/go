// 1295A. Display The Number
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
		if n%2 == 1 {
			fmt.Print(7)
			n -= 3
		}
		for n > 0 {
			fmt.Print(1)
			n -= 2
		}
		fmt.Println()
	}
}
