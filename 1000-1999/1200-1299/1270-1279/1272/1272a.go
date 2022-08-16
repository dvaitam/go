// 1272A. Three Friends
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
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		mi := min(min(a, b), c)
		mx := max(max(a, b), c)
		if mx-mi <= 2 {
			fmt.Println(0)
		} else {
			fmt.Println((mx - mi - 2) * 2)
		}
	}
}
