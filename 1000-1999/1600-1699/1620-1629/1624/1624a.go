// 1624A. Plus One on the Subset
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, n, curr int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &n)
		max := 0
		min := 1000000000
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			if curr < min {
				min = curr
			}
			if curr > max {
				max = curr
			}
		}
		fmt.Println(max - min)
	}
}
