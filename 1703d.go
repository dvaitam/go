// 1703D. Double Strings
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]string, n)
		m := make(map[string]bool)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = true
		}
		ans := make([]string, n)
		for i := 0; i < n; i++ {
			exists := false
			for j := 1; j < len(a[i]); j++ {
				if m[a[i][:j]] && m[a[i][j:]] {
					exists = true
					break
				}
			}
			if exists {
				ans[i] = "1"
			} else {
				ans[i] = "0"
			}
		}
		fmt.Println(strings.Join(ans, ""))
	}
}
