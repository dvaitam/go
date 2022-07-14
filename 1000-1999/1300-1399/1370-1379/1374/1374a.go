// 1374A. Required Remainder
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
		var x, y, n int
		fmt.Fscan(reader, &x, &y, &n)
		ans := (n/x)*x + y
		if ans > n {
			ans -= x
		}
		fmt.Println(ans)
	}
}
