// 281A. Word Capitalization
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text[0] >= 97 {
		fmt.Println(replaceAtIndex(text, text[0]-32, 0))
	} else {
		fmt.Println(text)
	}

}

func replaceAtIndex(in string, r byte, i int) string {
	out := []byte(in)
	out[i] = r
	return string(out)
}
