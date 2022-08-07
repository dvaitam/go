// 1283B. Candies Division
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
		rem := n % k
		ans := n - rem
		if rem > k/2 {
			rem = k / 2
		}
		fmt.Println(ans + rem)
	}
}
