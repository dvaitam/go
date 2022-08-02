// 1238A. Prime Subtraction
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
		var x, y int64
		fmt.Fscan(reader, &x, &y)
		if x-y == 1 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
