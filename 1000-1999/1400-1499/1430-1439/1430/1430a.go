// 1430A. Number of Apartments
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
		if n == 1 || n == 2 || n == 4 {
			fmt.Println(-1)
		} else {
			if n%3 == 0 {
				fmt.Println(n/3, 0, 0)
			} else if n%3 == 1 {
				fmt.Println((n-7)/3, 0, 1)
			} else {
				fmt.Println((n-5)/3, 1, 0)
			}
		}
	}
}
