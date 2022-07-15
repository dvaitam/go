// 1401A. Distance and Axis
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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		if k > n {
			fmt.Println(k - n)
		} else {
			if (k+n)%2 == 1 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
}
