// 1358A. Park Lighting
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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		ans := (n * m) / 2
		if (n*m)%2 == 1 {
			ans++
		}
		fmt.Println(ans)
	}
}
