// 1454A. Special Permutation
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
		fmt.Fscan(reader, &n)
		for i := 2; i <= n; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println(1)
	}
}
