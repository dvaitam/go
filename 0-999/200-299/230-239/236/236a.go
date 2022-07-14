// 236A. Boy or Girl
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make([]bool, 26, 26)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	unique := 0
	for i := 0; i < len(text); i++ {
		if counts[text[i]-'a'] == false {
			counts[text[i]-'a'] = true
			unique++
		}
	}
	if unique%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
