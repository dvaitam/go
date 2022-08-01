// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		zeros := 0
		ans := len(s)
		n := len(s)
		for i := n - 1; i >= 0; i-- {
			if s[i] == '0' {
				zeros++
			}
			if zeros == 2 {
				ans = min(ans, n-i-2)
				break
			}
			if zeros == 1 && s[i] == '5' {
				ans = min(ans, n-i-2)
				break
			}
		}
		found_five := false
		for i := n - 1; i >= 0; i-- {
			if s[i] == '5' {
				found_five = true
			}
			if found_five && (s[i] == '2' || s[i] == '7') {
				ans = min(ans, n-i-2)
				break
			}
		}
		fmt.Println(ans)
	}
}
