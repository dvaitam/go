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
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 1; i <= n; i++ {
			a[i-1] = i
		}
		for i := n - 1; i > 0; i -= 2 {
			a[i], a[i-1] = a[i-1], a[i]
		}
		for i := 0; i < n; i++ {
			fmt.Print(a[i])
			if i != n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
