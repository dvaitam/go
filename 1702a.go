// 1702A. Round Down the Price
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
		var m int
		fmt.Fscan(reader, &m)
		curr := 1
		for m/curr >= 10 {
			curr = curr * 10
		}
		fmt.Println(m - curr)
	}
}
