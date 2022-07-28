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
		var x, n int64
		fmt.Fscan(reader, &x, &n)
		if n%4 == 0 {
			fmt.Println(x)
		} else {
			ex := ((n - 1) / 4) * 4
			ex++
			if (n-1)%4 == 1 {
				ex -= n
			} else if (n-1)%4 == 2 {
				ex -= n
				ex -= n - 1
			} else if (n-1)%4 == 3 {
				ex -= n - 2
			}
			if x%2 == 0 {
				ex = x - ex
			} else {
				ex += x
			}
			fmt.Println(ex)
		}

	}
}
