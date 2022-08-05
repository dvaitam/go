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
		var n int
		fmt.Fscan(reader, &n)
		if n == 1 {
			fmt.Println(2)
		} else {
			ans := n / 3
			if n%3 != 0 {
				ans++
			}
			fmt.Println(ans)
		}
	}
}
