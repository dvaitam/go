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
		var s string
		fmt.Fscan(reader, &n, &s)
		if n%2 == 1 {
			one := false
			for i := 0; i < n; i += 2 {
				if (s[i]-'0')%2 == 1 {
					one = true
					break
				}
			}
			if one {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		} else {
			two := false
			for i := 1; i < n; i += 2 {
				if (s[i]-'0')%2 == 0 {
					two = true
					break
				}
			}
			if two {
				fmt.Println(2)
			} else {
				fmt.Println(1)
			}
		}
	}
}
