// 707A. Brain's Photos
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m, n int
	fmt.Fscan(reader, &m, &n)

	var curr string
	var mp = map[string]bool{"C": true, "M": true, "Y": true}
	ans := "#Black&White"
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &curr)
			if mp[curr] {
				ans = "#Color"
			}
		}
	}
	fmt.Println(ans)
}
