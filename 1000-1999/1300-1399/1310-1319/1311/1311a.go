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
		if a < b {
			if (b-a)%2 == 0 {
				fmt.Println(2)
			} else {
				fmt.Println(1)
			}
		} else if a > b {
			if (a-b)%2 == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		} else {
			fmt.Println(0)
		}
	}
}
