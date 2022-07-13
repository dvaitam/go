// 1348A. Phoenix and Balance
package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		mx := n/2 - 1
		half := 0
		total := 0
		for i := 1; i <= n; i++ {
			if i <= mx || i == n {
				half += 1 << i
			}
			total += 1 << i
		}
		//fmt.Println(mx, half, total)
		fmt.Println(abs(total - 2*half))
	}
}
