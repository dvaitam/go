// 00
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
	l := 0
	for n > 0 {
		r := (l + 1) * (l + 2) / 2
		if n >= r {
			l++
			n -= r
		} else {
			break
		}
	}
	fmt.Println(l)
}
