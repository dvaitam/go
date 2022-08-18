// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	curr := 1
	for n > 0 {
		curr *= 2
		n--
		if curr > m {
			break
		}
	}
	fmt.Println(m % curr)
}
