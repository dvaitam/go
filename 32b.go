// 32B. Borze
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &s)
	l := make([]byte, 0)
	i := 0
	for i < len(s) {
		if s[i] == '.' {
			l = append(l, '0')
			i++
		} else {
			if s[i+1] == '.' {
				l = append(l, '1')
			} else {
				l = append(l, '2')
			}
			i += 2
		}
	}
	fmt.Println(string(l))
}
