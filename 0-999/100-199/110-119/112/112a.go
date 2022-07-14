// 112A. Petya and Strings
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text1, _ := reader.ReadString('\n')
	text1 = strings.TrimSpace(text1)
	text1 = strings.ToLower(text1)
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	text2 = strings.ToLower(text2)
	if text1 < text2 {
		fmt.Println(-1)
	} else if text1 > text2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
