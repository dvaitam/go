// 282A - Bit++
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	n, _ = strconv.Atoi(text)

	ans := 0
	for i := 0; i < n; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "X++" || text == "++X" {
			ans++
		} else if text == "X--" || text == "--X" {
			ans--
		}

	}
	fmt.Println(ans)
}
