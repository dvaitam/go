// 344A. Magnets
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
	var text, prev string
	ans := 1
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(tmp))

	for i := 0; i < n; i++ {
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if i > 0 {
			if prev[1] == text[0] {
				ans++
			}
		}
		prev = text
	}

	fmt.Println(ans)
}
