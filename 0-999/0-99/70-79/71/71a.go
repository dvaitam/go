// 71A - Way Too Long Words
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
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			i--
			continue
		}

		if len(text) > 10 {
			fmt.Println(string(text[0]) + strconv.Itoa(len(text)-2) + string(text[len(text)-1]))
		} else {
			fmt.Println(text)
		}
	}
}
