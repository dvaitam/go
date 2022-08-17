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
		var n, m int64
		fmt.Fscan(reader, &n, &m)
		total := n + m - 1
		if total%2 == 1 {
			fmt.Println("Tonya")
		} else {
			fmt.Println("Burenka")
		}

	}
}
