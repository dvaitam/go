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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		if k%4 == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			if k%2 == 1 {
				for i := 1; i <= n; i += 2 {
					fmt.Println(i, i+1)
				}
			} else {
				for i := 1; i <= n; i += 2 {
					if (i+1)%4 == 2 {
						fmt.Println(i+1, i)
					} else {
						fmt.Println(i, i+1)
					}
				}
			}
		}

	}
}
