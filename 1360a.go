// 1360A. Minimal Square
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	var T, a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &a, &b)
		s := max(max(a, b), 2*min(a, b))
		fmt.Println(s * s)

	}
}
