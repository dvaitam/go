// 1213A. Chips Moving
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, curr int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	even := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &curr)
		if curr%2 == 0 {
			even++
		}
	}
	fmt.Println(min(even, n-even))
}
