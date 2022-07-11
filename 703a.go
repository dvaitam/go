// 703A. Mishka and Game
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	misk := 0
	chris := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a, &b)
		if a > b {
			misk++
		} else if b > a {
			chris++
		}
	}
	if misk > chris {
		fmt.Println("Mishka")
	} else if chris > misk {
		fmt.Println("Chris")
	} else {
		fmt.Println("Friendship is magic!^^")
	}
}
