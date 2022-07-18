// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		c := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &c[i])
		}
		even := make(map[int]int)
		odd := make(map[int]int)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				even[c[i]] = odd[c[i]] + 1
			} else {
				odd[c[i]] = even[c[i]] + 1
			}
		}

		for i := 1; i <= n; i++ {
			fmt.Print(max(even[i], odd[i]), " ")
		}
		fmt.Println()
	}
}
