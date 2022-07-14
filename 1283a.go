// 1283A. Minutes Before the New Year
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
		var h, m int
		fmt.Fscan(reader, &h, &m)
		ans := 0
		if m != 0 {
			ans += (60 - m)
			h++
		}
		ans += (24 - h) * 60
		fmt.Println(ans)
	}
}
