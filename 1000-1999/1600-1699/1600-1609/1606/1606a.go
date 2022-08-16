// 00
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
		var s string
		fmt.Fscan(reader, &s)
		sb := []byte(s)
		if sb[0] != sb[len(s)-1] {
			sb[0] = sb[len(s)-1]
		}
		fmt.Println(string(sb))
	}
}
