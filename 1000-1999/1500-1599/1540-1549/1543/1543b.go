// 1543B. Customising the Track
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
		var n, curr, rem int64
		fmt.Fscan(reader, &n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(reader, &curr)
			rem += curr
			rem %= n
		}
		fmt.Println(rem * (n - rem))
	}
}
