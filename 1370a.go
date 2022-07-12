// 1370A. Maximum GCD
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)

	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &n)
		if n%2 == 1 {
			n--
		}
		fmt.Println(n / 2)
	}
}
