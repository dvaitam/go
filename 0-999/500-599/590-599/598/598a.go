// 598A. Tricky Sum
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
		var n int64
		fmt.Fscan(reader, &n)
		ans := (n * (n + 1)) / 2
		curr := int64(1)
		for curr <= n {
			ans -= 2 * curr
			curr = curr * 2
		}
		fmt.Println(ans)
	}
}
