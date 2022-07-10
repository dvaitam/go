// 1702B. Polycarp Writes a String from Memory
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
		m := make(map[byte]int)
		ans := 0
		for i := 0; i < len(s); i++ {
			m[s[i]]++
			if len(m) == 4 {
				ans++
				m = make(map[byte]int)
				m[s[i]]++
			}
		}
		if len(m) > 0 {
			ans++
		}
		fmt.Println(ans)

	}
}
