// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m := map[byte]int{'Q': 9, 'R': 5, 'B': 3, 'N': 3, 'P': 1, 'K': 0, 'q': -9, 'r': -5, 'b': -3, 'n': -3, 'p': -1, 'k': 0}
	ans := 0
	for t := 1; t <= 8; t++ {
		var s string
		fmt.Fscan(reader, &s)
		for i := 0; i < len(s); i++ {
			if s[i] != '.' {
				ans += m[s[i]]
			}
		}
	}
	if ans > 0 {
		fmt.Println("White")
	} else if ans < 0 {
		fmt.Println("Black")
	} else {
		fmt.Println("Draw")
	}
}
