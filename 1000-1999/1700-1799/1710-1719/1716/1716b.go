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
		fmt.Println(n)
		for i := 1; i <= n; i++ {
			fmt.Fprint(os.Stdout, i)
			if i != n {
				fmt.Fprint(os.Stdout, " ")
			}
		}
		fmt.Fprint(os.Stdout, "\n")
		for k := n - 1; k >= 1; k-- {
			for i := 1; i <= n; i++ {
				if i != k {
					fmt.Fprint(os.Stdout, i)
					fmt.Fprint(os.Stdout, " ")
				}
			}
			fmt.Fprintln(os.Stdout, k)
		}
	}
}
