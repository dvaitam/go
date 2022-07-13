// 1353A. Most Unstable Array
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
		var m, n int
		fmt.Fscan(reader, &n, &m)
		if n == 1 {
			fmt.Println(0)
		} else if n == 2 {
			fmt.Println(m)
		} else {
			fmt.Println(2 * m)
		}
	}
}
