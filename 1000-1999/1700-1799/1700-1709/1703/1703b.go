// 1703B. ICPC Balloons
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
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		m := make(map[byte]bool)
		ans := 0
		for i := 0; i < n; i++ {
			if !m[s[i]] {
				m[s[i]] = true
				ans += 2
			} else {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
