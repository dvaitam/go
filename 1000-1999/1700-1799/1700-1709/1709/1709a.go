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
		var x int
		var a [3]int
		fmt.Fscan(reader, &x, &a[0], &a[1], &a[2])
		if a[0] == 1 || a[1] == 2 || a[2] == 3 || a[x-1] == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
