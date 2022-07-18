// 1555A. PizzaForces
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
		var n int64
		fmt.Fscan(reader, &n)
		if n < 6 {
			fmt.Println(15)
		} else {
			if n%2 == 1 {
				n++
			}
			fmt.Println(5 * n / 2)
		}
	}
}
