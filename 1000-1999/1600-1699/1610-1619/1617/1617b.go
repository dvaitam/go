// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a int, b int) int {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		for i := 2; i > 0; i++ {
			if gcd(n-1-i, i) == 1 {
				fmt.Println(i, n-1-i, 1)
				break
			}
		}
	}
}
