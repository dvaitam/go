// 1537B. Bad Boy
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
		var n, m, i, j int
		fmt.Fscan(reader, &n, &m, &i, &j)
		fmt.Println(1, 1, n, m)
	}
}
