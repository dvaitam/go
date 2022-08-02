// 1714B. Remove Prefix
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
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		m := make(map[int]bool)
		ans := 0
		for i := n - 1; i >= 0; i-- {
			if m[a[i]] {
				ans = i + 1
				break
			} else {
				m[a[i]] = true
			}
		}
		fmt.Println(ans)
	}
}
