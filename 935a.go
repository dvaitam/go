// 935A. Fafa and his Company
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	ans := 0
	for l := 1; l <= n/2; l++ {
		if n%l == 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
