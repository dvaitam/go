// 1426A. Floor Number
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
		var n, x int
		fmt.Fscan(reader, &n, &x)
		if n <= 2 {
			fmt.Println(1)
		} else {
			if (n-2)%x == 0 {
				fmt.Println((n-2)/x + 1)
			} else {
				fmt.Println((n-2)/x + 2)
			}
		}

	}
}
