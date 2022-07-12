// 749A. Bachgold Problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	fmt.Println(n / 2)
	if n%2 == 0 {
		fmt.Print(2)
	} else {
		fmt.Print(3)
	}
	for i := 1; i < n/2; i++ {
		fmt.Print(" ", 2)
	}
	fmt.Println()
}
