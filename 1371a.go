// 1371A. Magical Sticks
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
		ans := n / 2
		if n%2 == 1 {
			ans++
		}
		fmt.Println(ans)

	}
}
