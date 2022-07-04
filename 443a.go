// 443A. Anton and Letters
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			m[s[i]] = true
		}
	}
	fmt.Println(len(m))
}
