// 1433A. Boring Apartments
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var T, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &n)
		d := n % 10
		ans := (d - 1) * 10
		l := strconv.Itoa(n)
		if len(l) == 4 {
			ans += 10
		} else if len(l) == 3 {
			ans += 6
		} else if len(l) == 2 {
			ans += 3
		} else {
			ans += 1
		}
		fmt.Println(ans)
	}
}
