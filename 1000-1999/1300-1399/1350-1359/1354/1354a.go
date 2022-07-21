// 1354A. Alarm Clock
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
		var a, b, c, d int64
		fmt.Fscan(reader, &a, &b, &c, &d)
		if a <= b {
			fmt.Println(b)
		} else {
			if d >= c {
				fmt.Println(-1)
			} else {
				diff := c - d
				ans := b
				if (a-b)%diff == 0 {
					ans += ((a - b) / diff) * (c)
				} else {
					ans += ((a-b)/diff)*(c) + c
				}
				fmt.Println(ans)
			}
		}
	}
}
